package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	"encoding/hex"
	"encoding/json"

	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/glebarez/sqlite"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/didip/tollbooth_gin"
)

var DB *gorm.DB
var secretKey []byte
var brainData map[string][]float64
var geniusData map[string][]float64
var averageData map[string][]float64
var stupidData map[string][]float64

type Data struct {
    OTP          string       `json:"otp"`        
    Times        int          `json:"times"`
    Action       string       `json:"action"`
    Nickname     string       `json:"nickname"`
    Email        string       `json:"email"`
    Password     string       `json:"password"`
    Token        string       `json:"token"`
	FirstName    string       `json:"firstname"`
	Lastname     string       `json:"lastname"`
	Birthdate    string       `json:"birthdate"`
    Gender       string       `json:"gender"`
    Device       DeviceInfo   `json:"device"`
    Network      NetworkInfo  `json:"network"`
}

type DeviceInfo struct {
    ScreenWidth  int    `json:"screen_w"`
    ScreenHeight int    `json:"screen_h"`
    OS           string `json:"os"`
    Browser      string `json:"browser"`
    CPUCores     string `json:"cpu_cores"`
    Lang         string `json:"lang"`
    Timezone     string `json:"timezone"`
    PrefersDark  string   `json:"prefers_dark"`
}

type NetworkInfo struct {
    IP      string `json:"ip"`
    Country string `json:"country"`
    City    string `json:"city"`
    ISP     string `json:"isp"`
}

type User struct {
	gorm.Model
	Nickname     string       
	Email        string
	Password     string
	ScreenWidth  string
	ScreenHeight string
	OS           string
	Browser      string
	CPUCores     string
	TokenHash    string 
	NicknameHash string
	Wins     	 string  
	Losses   	 string  
	Draws    	 string  
	Total    	 string  
	MatchIDs     string
}

type OTPVerification struct {
	gorm.Model
	Email        string
	OTP          string
	NicknameHash string `gorm:"uniqueIndex"`
	Action       string
	Nickname     string
	Password     string
	ScreenWidth  string
	ScreenHeight string
	OS           string
	Browser      string
	CPUCores     string
	Times        string
	Token        string
}

type GameRequest struct {
    Token      string   `json:"token"`
    Board      []string `json:"board"`
    Click      int      `json:"click"`
    Difficulty string   `json:"difficulty"`
    WhoStarts  string   `json:"who_starts"`
    Symbol     string   `json:"symbol"`
    Timer      string   `json:"timer"`
}

type Games struct {
	gorm.Model
	GameHash      string `gorm:"uniqueIndex"`
	UserID        string
	WhoStarts     string
	IsGiveUp      string
	IsTimeOut     string
	Difficulty    string
	PlayerSymbol  string
	Result        string
	Board         string
	Move0         string
	Move1         string
	Move2         string
	Move3         string
	Move4         string
	Move5         string
	Move6         string
	Move7         string
	Move8         string
	ComputerMoves string

}

func Encrypt(text string) string {
	block, _ := aes.NewCipher(secretKey)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) string {
	data, err := hex.DecodeString(cryptoText)
	if err != nil {
		return ""
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return ""
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return ""
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return ""
	}
	return string(plaintext)
}


func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("XOkingdom.db"), &gorm.Config{})
	if err != nil {
		panic("connection error :" + err.Error())
	}
	err = DB.AutoMigrate(&User{}, &OTPVerification{}, &Games{})
	if err != nil {
		panic("error :" + err.Error())
	}
}

func generateOTP() int {
	max := big.NewInt(900000)
	num, _ := rand.Int(rand.Reader, max)
	return int(num.Int64()) + 100000
}

func is_malicious_game(request GameRequest) bool {
	if request.Difficulty != "hard" && request.Difficulty != "medium" && request.Difficulty != "easy" {
		return true
	}else if request.WhoStarts != "player" && request.WhoStarts != "computer" {
		return true
	}else if request.Symbol != "x" && request.Symbol != "o" {
		return true
	}else if request.Click < 0 || request.Click > 8 {
		return true
	}else if request.Timer != "10seconds" && request.Timer != "30seconds" && request.Timer != "1minute" {
		return true
	}else if len(request.Board) != 9 {
		return true
	}
	for _, cell := range request.Board {
		if cell != "" && cell != "X" && cell != "O" {
			return true
		}	
	}
	countX := 0
	countO := 0
	count  := 0
	for _, cell := range request.Board {
		if cell == "X" {
			countX++ 
			count++
		}
		if cell == "O" {
			countO++
			count++
		}
		if cell == "" {
			count++
		}
	}
	if countX - countO != 1 && countX - countO != 0 && countX - countO != -1 || count != 9{
		return true
	}

	return false
}

func loadBrain() map[string][]float64 {
	file, _ := os.Open("brain.json")
	defer file.Close()
	var brain map[string][]float64
	json.NewDecoder(file).Decode(&brain)
	return brain
}

func bestMove(board []string, brain map[string][]float64) int {
	state := ""
	for _, cell := range board {
		switch cell {
		case "X":
			state += "1"
		case "O":
			state += "2"
		default:
			state += "0"
		}
	}

		qValues := brain[state]
	bestMove, bestQ := -1, -1e9
	for i, cell := range board {
		if cell == "" {
			if qValues != nil && qValues[i] > bestQ {
				bestQ = qValues[i]
				bestMove = i
			} else if qValues == nil && bestMove == -1 {
				bestMove = i
			}
		}
	}

	return bestMove
}

func is_win (board []string)string {
	if board[0] != "" && board[0] == board[1] && board[0] == board[2] {
		return board[0]
	}
	if board[3] != "" && board[3] == board[4] && board[3] == board[5] {
		return board[3]
	}
	if board[6] != "" && board[6] == board[7] && board[6] == board[8] {
		return board[6]
	}
	if board[0] != "" && board[0] == board[3] && board[0] == board[6] {
		return board[0]
	}
	if board[1] != "" && board[1] == board[4] && board[1] == board[7] {
		return board[1]
	}
	if board[2] != "" && board[2] == board[5] && board[2] == board[8] {
		return board[2]
	}
	if board[0] != "" && board[0] == board[4] && board[0] == board[8] {
		return board[0]
	}
	if board[2] != "" && board[2] == board[4] && board[2] == board[6] {
		return board[2]
	}
	return "nil"
}

func brain_to_display(brain string) []string {
    display := make([]string, 9)
    for i, c := range brain {
        switch c {
        case '1': display[i] = "X"
        case '2': display[i] = "O"
        }
    }
    return display
}

func current_board(board []string) (string, []string) {
    brain := ""
    for _, cell := range board {
        switch cell {
        case "X", "x": brain += "1"
		case "O", "o": brain += "2"
        default: brain += "0"
        }
    }
    return brain, board
}

func give_up(c *gin.Context) {
	var request GameRequest
	var user User
	var game Games
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	tokenHash := HashString(request.Token)
	if err := DB.Where("token_hash = ?", tokenHash).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	
	ids := strings.Split(user.MatchIDs, ",")
	lastID := Decrypt(ids[len(ids)-1])
	if ids[0] == "" {
		switch request.Difficulty {
			case "hard":
				brainData = geniusData
			case "medium":
				brainData = averageData
			case "easy":
				brainData = stupidData
			default:
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
		}
		bestmove := bestMove(request.Board, brainData)
		var newGame Games
		gameHash := Encrypt(user.NicknameHash) 
			newGame = Games{
				GameHash: gameHash,
				UserID:    Encrypt(user.NicknameHash),
				WhoStarts: Encrypt("computer"),
				Difficulty: Encrypt(request.Difficulty),
				PlayerSymbol: Encrypt(request.Symbol),
				Move0: Encrypt(strconv.Itoa(bestmove)),
				ComputerMoves:Encrypt(strconv.Itoa(bestmove)),
				IsTimeOut: Encrypt("false"),
				IsGiveUp: Encrypt("true"),
				Result: Encrypt("computerwin"),

			}
			if user.MatchIDs == "" {
				user.MatchIDs = Encrypt(gameHash)
			} else {
				user.MatchIDs = strings.Join(append(ids, Encrypt(gameHash)), ",")
			}
			DB.Create(&newGame)
			DB.Save(&user)
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
			return

	}

	if err := DB.Where("game_hash = ?", lastID).First(&game).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	game.IsGiveUp = Encrypt("true")
	game.Result = Encrypt("computerwin")
	DB.Save(&game)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func time_out(c *gin.Context) {
	var request GameRequest
	var user User
	var game Games
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	tokenHash := HashString(request.Token)
	if err := DB.Where("token_hash = ?", tokenHash).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	
	ids := strings.Split(user.MatchIDs, ",")
	lastID := Decrypt(ids[len(ids)-1])
	if ids[0] == "" {
		switch request.Difficulty {
			case "hard":
				brainData = geniusData
			case "medium":
				brainData = averageData
			case "easy":
				brainData = stupidData
			default:
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
		}
		bestmove := bestMove(request.Board, brainData)
		var newGame Games
		gameHash := Encrypt(user.NicknameHash) 
			newGame = Games{
				GameHash: gameHash,
				UserID:    Encrypt(user.NicknameHash),
				WhoStarts: Encrypt("computer"),
				IsGiveUp: Encrypt("false"),
				Difficulty: Encrypt(request.Difficulty),
				PlayerSymbol: Encrypt(request.Symbol),
				Move0: Encrypt(strconv.Itoa(bestmove)),
				ComputerMoves:Encrypt(strconv.Itoa(bestmove)),
				IsTimeOut: Encrypt("true"),
				Result: Encrypt("computerwin"),
				
			}
			if user.MatchIDs == "" {
				user.MatchIDs = Encrypt(gameHash)
			} else {
				user.MatchIDs = strings.Join(append(ids, Encrypt(gameHash)), ",")
			}
			DB.Create(&newGame)
			DB.Save(&user)
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
			return
	}
	if err := DB.Where("game_hash = ?", lastID).First(&game).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	game.IsTimeOut = Encrypt("true")
	game.Result = Encrypt("computerwin")
	DB.Save(&game)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})

	
}

func is_malicious_move(stored, request []string, player_upper string) bool {
    changes := 0
    for i := 0; i < 9; i++ {
        if stored[i] != request[i] {
            changes++
            if request[i] != player_upper || stored[i] != "" {
                return true
            }
        }
    }
    return changes > 1 
}

func make_move(c *gin.Context) {
	var request GameRequest
	var user User
	var game Games
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	tokenHash := HashString(request.Token)

	if err := DB.Where("token_hash = ?", tokenHash).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	ids := strings.Split(user.MatchIDs, ",")
	lastID := Decrypt(ids[len(ids)-1])
	if err := DB.Where("game_hash = ?", lastID).First(&game).Error; err != nil && lastID != "" {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	old_display := brain_to_display(Decrypt(game.Board))

	playerSymbol := "X"
	if request.Symbol == "o" {
		playerSymbol = "O"
	}
	
	if is_malicious_game(request) {
		c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
		return
	}
	var roundnum int
	for _, rounds := range request.Board {
		if rounds == "X" || rounds == "O" {
			roundnum++
		}
	}
	if roundnum != 0 && lastID != "" && Decrypt(game.Result) == "running" {
		if is_malicious_move(old_display, request.Board, playerSymbol) {
			c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
			return
		}
	}

		computerSymbol := "O"
			if request.Symbol == "o" {
				computerSymbol = "X"
			}


			
		

	var newGame Games
	if roundnum == 0 {
		switch request.Difficulty {
			case "hard":
				brainData = geniusData
			case "medium":
				brainData = averageData
			case "easy":
				brainData = stupidData
			default:
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
		}
		if Decrypt(game.Result) == "computerwin" || Decrypt(game.Result) == "userwin" || Decrypt(game.Result) == "draw" || Decrypt(game.Result) == "" && lastID == "" {
			bestmove := bestMove(request.Board, brainData)
			
			if is_win(request.Board) != "nil" || request.WhoStarts != "computer" {
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
			}
			
			request.Board[bestmove] = computerSymbol
			board01, boardxo := current_board(request.Board)
			

			

			gameHash := Encrypt(user.NicknameHash) 
			newGame = Games{
				GameHash: gameHash,
				UserID:    Encrypt(user.NicknameHash),
				WhoStarts: Encrypt("computer"),
				IsGiveUp: Encrypt("false"),
				Difficulty: Encrypt(request.Difficulty),
				PlayerSymbol: Encrypt(request.Symbol),
				Result: Encrypt("running"),
				Move0: Encrypt(strconv.Itoa(bestmove)),
				ComputerMoves:Encrypt(strconv.Itoa(bestmove)),

			}
			if user.MatchIDs == "" {
				user.MatchIDs = Encrypt(gameHash)
			} else {
				user.MatchIDs = strings.Join(append(ids, Encrypt(gameHash)), ",")
			}
			if err := DB.Create(&newGame).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
				return
			}
			DB.Save(&user)
			
			newGame.Board = Encrypt(board01)
			
			DB.Save(&newGame)
			if is_win(boardxo) != "nil" {
				c.JSON(http.StatusOK, gin.H{"status": "LOSE", "board": boardxo})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "OK", "board": boardxo})
			return

		}
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	} else if roundnum == 1 {
		switch request.Difficulty {
			case "hard":
				brainData = geniusData
			case "medium":
				brainData = averageData
			case "easy":
				brainData = stupidData
			default:
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
		}
		if Decrypt(game.Result) == "computerwin" || Decrypt(game.Result) == "userwin" || Decrypt(game.Result) == "draw" || Decrypt(game.Result) == "" && lastID == "" {
			bestmove := bestMove(request.Board, brainData)
			
			if is_win(request.Board) != "nil" || request.WhoStarts != "player" {
				c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
				return
			}
			
			request.Board[bestmove] = computerSymbol
			board01, boardxo := current_board(request.Board)
			

			

			gameHash := Encrypt(user.NicknameHash) 
			newGame = Games{
				GameHash: gameHash,
				UserID:    Encrypt(user.NicknameHash),
				WhoStarts: Encrypt("player"),
				IsGiveUp: Encrypt("false"),
				Difficulty: Encrypt(request.Difficulty),
				PlayerSymbol: Encrypt(request.Symbol),
				Result: Encrypt("running"),
				Move0: Encrypt(strconv.Itoa(request.Click)),
				Move1: Encrypt(strconv.Itoa(bestmove)),
				ComputerMoves:Encrypt(strconv.Itoa(bestmove)),

			}
			if user.MatchIDs == "" {
				user.MatchIDs = Encrypt(gameHash)
			} else {
				user.MatchIDs = strings.Join(append(ids, Encrypt(gameHash)), ",")
			}
			if err := DB.Create(&newGame).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
				return
			}
			DB.Save(&user)
			
			newGame.Board = Encrypt(board01)
			
			
			DB.Save(&newGame)
			c.JSON(http.StatusOK, gin.H{"status": "OK", "board": boardxo})
			return

		}
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return

		} else if roundnum >= 2 && roundnum <= 7 {

			
			board01, boardxo := current_board(request.Board)

			game.Board = Encrypt(board01)
			

			switch roundnum {
			case 2: game.Move1 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move2 = Encrypt(strconv.Itoa(bestmove))
			case 3: game.Move2 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move3 = Encrypt(strconv.Itoa(bestmove))
			case 4: game.Move3 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move4 = Encrypt(strconv.Itoa(bestmove))
			case 5: game.Move4 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move5 = Encrypt(strconv.Itoa(bestmove))
			case 6: game.Move5 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move6 = Encrypt(strconv.Itoa(bestmove))
			case 7: game.Move6 = Encrypt(strconv.Itoa(request.Click))
					switch game.Difficulty {
						case "hard":
							brainData = geniusData
						case "medium":
							brainData = averageData
						case "easy":
							brainData = stupidData
						default:
							c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
							return
					}
					bestmove := bestMove(request.Board, brainData)
					game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
					request.Board[bestmove] = computerSymbol
					game.Move7 = Encrypt(strconv.Itoa(bestmove))
			}

			if is_win(boardxo) == computerSymbol {
				game.Result = Encrypt("computerwin")
				losses,_ := strconv.Atoi(Decrypt(user.Losses))
				losses++
				user.Losses = Encrypt(strconv.Itoa(losses))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "LOSE", "board": boardxo})
				return
			}
			if is_win(boardxo) == playerSymbol {
				game.Result = Encrypt("userwin")
				wins,_ := strconv.Atoi(Decrypt(user.Wins))
				wins++
				user.Wins = Encrypt(strconv.Itoa(wins))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "WIN", "board": boardxo})
				return
			}
			DB.Save(&game)
			c.JSON(http.StatusOK, gin.H{"status": "OK", "board": boardxo})
			return
	}else if roundnum == 8 {
		bestmove := bestMove(request.Board, brainData)
		request.Board[bestmove] = computerSymbol
		board01, boardxo := current_board(request.Board)
			
		game.Board = Encrypt(board01)
		game.Move7 = Encrypt(strconv.Itoa(request.Click))
		game.Move8 = Encrypt(strconv.Itoa(bestmove))
		game.ComputerMoves = Encrypt(Decrypt(game.ComputerMoves) + strconv.Itoa(bestmove))
		winner := is_win(boardxo)
		
		switch winner {
			case "nil":
				game.Result = Encrypt("draw")
				draws,_ := strconv.Atoi(Decrypt(user.Draws))
				draws++
				user.Draws = Encrypt(strconv.Itoa(draws))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "DRAW", "board": boardxo})
				return
			case computerSymbol:
				game.Result = Encrypt("computerwin")
				losses,_ := strconv.Atoi(Decrypt(user.Losses))
				losses++
				user.Losses = Encrypt(strconv.Itoa(losses))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "LOSE", "board": boardxo})
				return
				
			default:
				game.Result = Encrypt("userwin")
				wins,_ := strconv.Atoi(Decrypt(user.Wins))
				wins++
				user.Wins = Encrypt(strconv.Itoa(wins))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "WIN", "board": boardxo})
				return
			}
		}else if roundnum == 9 {
		board01, boardxo := current_board(request.Board)
			
		game.Board = Encrypt(board01)
		game.Move8 = Encrypt(strconv.Itoa(request.Click))
		winner := is_win(boardxo)
		
		switch winner {
			case "nil":
				game.Result = Encrypt("draw")
				draws,_ := strconv.Atoi(Decrypt(user.Draws))
				draws++
				user.Draws = Encrypt(strconv.Itoa(draws))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "DRAW", "board": boardxo})
				return
			case computerSymbol:
				game.Result = Encrypt("computerwin")
				losses,_ := strconv.Atoi(Decrypt(user.Losses))
				losses++
				user.Losses = Encrypt(strconv.Itoa(losses))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "LOSE", "board": boardxo})
				return
				
			default:
				game.Result = Encrypt("userwin")
				wins,_ := strconv.Atoi(Decrypt(user.Wins))
				wins++
				user.Wins = Encrypt(strconv.Itoa(wins))
				total,_ := strconv.Atoi(Decrypt(user.Total))
				total++
				user.Total = Encrypt(strconv.Itoa(total))
				DB.Save(&game)
				DB.Save(&user)
				c.JSON(http.StatusOK, gin.H{"status": "WIN", "board": boardxo})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
	
}

func resume_playing(c *gin.Context) {
	var request GameRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	tokenHash := HashString(request.Token)
	var user User
	if err := DB.Where("token_hash = ?", tokenHash).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	if user.MatchIDs == "" {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	ids := strings.Split(user.MatchIDs, ",")
	lastID := Decrypt(ids[len(ids)-1])
	var game Games
	if err := DB.Where("game_hash = ?", lastID).First(&game).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	if Decrypt(game.Result) == "computerwin" || Decrypt(game.Result) == "userwin" || Decrypt(game.Result) == "draw" {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	if Decrypt(game.Move0) == "" {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}

	stored_board := brain_to_display(Decrypt(game.Board))
	_, boardxo := current_board(stored_board)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "board": boardxo})
}

func is_malicious(request Data) bool {
	ismatched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", request.Nickname)
	if !ismatched || len(request.Nickname) > 15 || len(request.Nickname) < 3 {
		return true
	}
	if request.Device.ScreenWidth < 30 || request.Device.ScreenHeight < 30 {
		return true
	}
	if request.Action != "login" && request.Action != "signin" {
		return true
	}
	at := false
	point := false
	for _, letter := range request.Email {
		switch letter {
		case '@':
			at = true
		case '.':
			point = true
		}
	}
	if !at || !point {
		return true
	}
	if len(request.Password) < 8 {
		return true
	}
	if request.Action == "signin" {
		if request.Gender != "Male" && request.Gender != "Female" {
			return true
		} else {
			layout := "2006-01-02"
    		birth, err := time.Parse(layout, request.Birthdate)
			now := time.Now()
			maxDate := now.AddDate(-150, 0, 0)
			minDate := now.AddDate(-18, 0, 0)
			if err != nil {
        		return true
    		}else if birth.After(now){
				return true
			}else if birth.After(minDate) {
				return true
			}else if birth.Before(maxDate) {
    			return true       
			}
		}
	}
	if request.Times >= 5 {
		return true
	}
	return false

}

func HashString(text string) string {
	data := []byte(text)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	keyHex := os.Getenv("SECRET_KEY_HEX")
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic("Error decoding secret key: " + err.Error())
	}
	secretKey = key

	genius, err := os.Open("genius.json")
	if err != nil {
		panic("Error loading genius.json: " + err.Error())
	}
	defer genius.Close()
	json.NewDecoder(genius).Decode(&geniusData)

	stupid, err := os.Open("stupid.json")
	if err != nil {
		panic("Error loading stupid.json: " + err.Error())
	}
	defer stupid.Close()
	json.NewDecoder(stupid).Decode(&stupidData)

	average, err := os.Open("average.json")
	if err != nil {
		panic("Error loading average.json: " + err.Error())
	}
	defer average.Close()
	json.NewDecoder(average).Decode(&averageData)
}

func autologin(c *gin.Context) {
	var request struct {
        Token string `json:"token"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
        return
    }
	tokenHash := HashString(request.Token)
    var user User
    result := DB.Where("token_hash = ?", tokenHash).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
    } else {
        c.JSON(http.StatusOK, gin.H{"status": "OK"})
    }
}


func checkOTP(c *gin.Context) {
	var request Data
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	hash := HashString(request.Nickname)
	var db_data OTPVerification
	result := DB.Where("nickname_hash = ?", hash).First(&db_data)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	if Decrypt(db_data.OTP) != request.OTP || Decrypt(db_data.Email) != request.Email {
    timesInt, _ := strconv.Atoi(Decrypt(db_data.Times))
    if timesInt >= 5 {
        c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
        return
    }
    timesInt++
    db_data.Times = Encrypt(strconv.Itoa(timesInt))
    DB.Save(&db_data)
    c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
    return
	}
	var newUser User
	action := Decrypt(db_data.Action)
	switch action {
		case "signin":
			token := Encrypt(Encrypt(request.Nickname)) 
			tokenHash := HashString(token) 
			newUser = User{
				Nickname:     Decrypt(db_data.Nickname),     
				NicknameHash: hash,
				Email:        Encrypt(Decrypt(db_data.Email)),   
				Password:     Encrypt(Decrypt(db_data.Password)), 
				ScreenWidth:  Encrypt(Decrypt(db_data.ScreenWidth)),       
				ScreenHeight: Encrypt(Decrypt(db_data.ScreenHeight)),
				OS:           Encrypt(Decrypt(db_data.OS)),
				Browser:      Encrypt(Decrypt(db_data.Browser)),
				CPUCores:     Encrypt(Decrypt(db_data.CPUCores)),
				TokenHash:    tokenHash,
			}
			DB.Create(&newUser)
			DB.Unscoped().Delete(&db_data)
			c.JSON(http.StatusOK, gin.H{"status": "OK", "token": token})
		case "login":
				token := Encrypt(Encrypt(request.Nickname)) 
				tokenHash := HashString(token)
				var existingUser User
				result := DB.Where("nickname_hash = ?", hash).First(&existingUser)
				if result.Error != nil {
					c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
					return
				}
				if Decrypt(existingUser.Email) != request.Email {
					c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
					return
				}
			existingUser.TokenHash    = tokenHash
			existingUser.NicknameHash = hash
			DB.Save(&existingUser)
			c.JSON(http.StatusOK, gin.H{"status": "OK", "token": token})
		default:
			c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
			return
	}
		
	
	}

func catchfirstrequest(c *gin.Context) {
	var request Data
	DB.Unscoped().Where("created_at < ?", time.Now().Add(-10*time.Minute)).Delete(&OTPVerification{})
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
		return
	}
	if is_malicious(request) {
		c.JSON(http.StatusOK, gin.H{"status": "MALICIOUS"})
		return
	}
    if request.Action == "login" {
        var existingUser User
        result := DB.Where("nickname_hash = ?", HashString(request.Nickname)).First(&existingUser)
        if result.Error != nil {
            c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
            return
        }
        decryptedPassword := Decrypt(existingUser.Password)
		decryptedEmail := Decrypt(existingUser.Email)

		if decryptedPassword != request.Password || decryptedEmail != request.Email {
			c.JSON(http.StatusOK, gin.H{"status": "WRONG"})
			return
		}
    }
	otp := generateOTP()
	newEntry := OTPVerification{
		NicknameHash: HashString(request.Nickname),
		Action:       Encrypt(request.Action),
		Email:        Encrypt(request.Email),
		OTP:          Encrypt(strconv.Itoa(otp)),
		Nickname:     Encrypt(request.Nickname),
		Password:     Encrypt(request.Password),
		ScreenWidth:  Encrypt(strconv.Itoa(request.Device.ScreenWidth)),
		ScreenHeight: Encrypt(strconv.Itoa(request.Device.ScreenHeight)),
		OS:           Encrypt(request.Device.OS),
		Browser:      Encrypt(request.Device.Browser),
		CPUCores:     Encrypt(request.Device.CPUCores),
		Times:        Encrypt("0"),
		Token:        Encrypt(request.Nickname),
	}
	DB.Unscoped().Where("nickname_hash = ?", HashString(request.Nickname)).Delete(&OTPVerification{})
	DB.Create(&newEntry)
	message := gomail.NewMessage()
	smtpUser := os.Getenv("EMAIL_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	message.SetHeader("From", smtpUser)
	message.SetHeader("To", request.Email)
	htmlBody := fmt.Sprintf(`
							<!DOCTYPE html>
							<html>
							<head>
								<meta charset="UTF-8">
								<meta name="viewport" content="width=device-width, initial-scale=1.0">
								<meta name="color-scheme" content="light dark">
								<meta name="supported-color-schemes" content="light dark">
								<style>
									:root {
										color-scheme: light dark;
										supported-color-schemes: light dark;
									}
									@media (prefers-color-scheme: dark) {
										body, .bg-main { background-color: #1a0a2e !important; }
										.bg-card { background-color: #101024 !important; }
										.bg-code { background-color: #0a0a1a !important; }
										.text-light { color: #f7ebff !important; }
										.text-purple { color: #7c3aed !important; }
									}
								</style>
							</head>
							<body style="margin:0;padding:0;background-color:#1a0a2e;font-family:Arial,sans-serif;" class="bg-main">
								<table width="100%%" cellpadding="0" cellspacing="0" style="background-color:#1a0a2e;padding:40px 0;" class="bg-main">
									<tr>
										<td align="center">
											<table width="480" cellpadding="0" cellspacing="0" style="background-color:#101024;border-radius:12px;overflow:hidden;" class="bg-card">
												<tr>
													<td style="padding:40px 30px;text-align:center;">
														<h1 style="color:#f7ebff;font-size:28px;margin:0 0 10px;" class="text-light">
															<span style="color:#0000ff;">X</span><span style="color:#800080;">O</span><span style="color:#f7ebff;" class="text-light">kingdom</span>
														</h1>
														<p style="color:#f7ebff;font-size:16px;margin:20px 0;" class="text-light">Your verification code is:</p>
														<p style="font-size:48px;letter-spacing:12px;color:#7c3aed !important;font-weight:bold;margin:20px 0;background-color:#0a0a1a;padding:16px;border-radius:8px;" class="bg-code text-purple">%d</p>
														<p style="color:#9e9eb8;font-size:13px;margin:20px 0 0;">This code expires in 10 minutes.</p>
													</td>
												</tr>
											</table>
										</td>
									</tr>
								</table>
							</body>
							</html>

							`, otp)

	message.SetBody("text/html", htmlBody)
	d := gomail.NewDialer("smtp-relay.brevo.com", 587, "ad0d4a001@smtp-brevo.com", smtpPass)
	if err := d.DialAndSend(message); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func main() {
	initDB()
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.SetTrustedProxies(nil)
	server.Use(cors.Default())

	gameLimiter := tollbooth.NewLimiter(5.0, &limiter.ExpirableOptions{
		DefaultExpirationTTL: 1 * time.Minute,
	})
	gameLimiter.SetBurst(8)

	otpLimiter := tollbooth.NewLimiter(1.0/60.0, &limiter.ExpirableOptions{
		DefaultExpirationTTL: 5 * time.Minute,
	})
	otpLimiter.SetBurst(1)

	server.POST("/api/auth/auto-login", autologin)
	server.POST("/api/auth/verify-otp", checkOTP)
	server.POST("/api/auth/send-otp", tollbooth_gin.LimitHandler(otpLimiter), catchfirstrequest)

	gameGroup := server.Group("/api/game/xo/3x3")
	gameGroup.Use(tollbooth_gin.LimitHandler(gameLimiter))
	{
		gameGroup.POST("/move", make_move)
		gameGroup.POST("/give-up", give_up)
		gameGroup.POST("/resume-playing", resume_playing)
		gameGroup.POST("/time-out", time_out)
	}

	server.Run(":8080")
}
