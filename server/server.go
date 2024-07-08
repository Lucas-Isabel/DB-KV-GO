package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lucasbyte/DB-KV-GO/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Store *storage.Storage
}

func NewServer() *Server {
	return &Server{Store: storage.NewStorage()}
}

type Request_Post struct {
	Method string `json:"metodo"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type Request_Get struct {
	Method string `json:"metodo"`
	Key    string `json:"key"`
}

func (s *Server) SETk(c *gin.Context) {
	var req Request_Post

	// Vinculando JSON recebido à estrutura Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Method == "SET" {
		// Configurando o valor no armazenamento
		valid := s.Store.SETk(req.Key, req.Value)
		if valid {
			c.JSON(http.StatusOK, gin.H{
				"message": "Value set successfully!",
				"key":     req.Key,
				"value":   req.Value,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unsupported method",
		})
	}
}

func (s *Server) GETk(c *gin.Context) {
	var req Request_Get

	// Vinculando JSON recebido à estrutura Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Method == "GET" {
		// Configurando o valor no armazenamento
		value, includes := s.Store.GETk(req.Key)
		if includes {
			c.JSON(http.StatusOK, gin.H{
				"message": "Value get successfully!",
				"value":   value,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unsupported method",
		})
	}
}

func (s *Server) ALLkv(c *gin.Context) {
	var req Request_Get

	// Vinculando JSON recebido à estrutura Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Method == "ALL" {
		// Configurando o valor no armazenamento
		values := s.Store.ALLkv()
		if len(values) > 0 {
			log.Println(values)

			// Abrir um arquivo para escrita
			file, err := os.Create("output.txt")
			if err != nil {
				fmt.Println("Erro ao criar o arquivo:", err)
				return
			}
			defer file.Close()

			// Escrever a struct Person no arquivo
			_, err = fmt.Fprintf(file, "%+v\n", values)
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo:", err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Values get successfully!",
				"values":  values,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unsupported method",
		})
	}
}

func (s *Server) DELETEk(c *gin.Context) {
	var req Request_Get

	// Vinculando JSON recebido à estrutura Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Method == "DEL" {
		// Configurando o valor no armazenamento
		deletes := s.Store.Delete(req.Key)
		if deletes {
			c.JSON(http.StatusOK, gin.H{
				"message": "Value DELETE successfully!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Value NOT exist!",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unsupported method",
		})
	}

}
