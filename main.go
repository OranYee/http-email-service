package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

// 邮件内容结构
type Email struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// SMTP 服务器配置
const (
	SMTPServer = "smtp.example.com"
	SMTPPort   = "25"
)

// 处理邮件发送请求的处理程序
func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	// 解码POST请求中的JSON数据
	var email Email
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 构建邮件消息
	msg := "From: " + email.From + "\r\n" +
		"To: " + strings.Join(email.To, ", ") + "\r\n" +
		"Subject: " + email.Subject + "\r\n\r\n" +
		email.Body

	// 使用SMTP发送电子邮件
	err = smtp.SendMail(SMTPServer+":"+SMTPPort, nil, email.From, email.To, []byte(msg))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

func main() {
	// 设置邮件发送处理程序
	http.HandleFunc("/send-email", sendEmailHandler)

	// 启动HTTP服务器
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
