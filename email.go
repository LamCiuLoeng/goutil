package goutil 

import (
    "net/smtp"
    "net/mail"
    "strings"
    "fmt"
    "encoding/base64"
)

func SendEmail(from string,
                 to []string,
		 cc []string,
		 subject string,
		 content string,
		 serverip string,
		 port string) error{

    emailfrom := mail.Address{from,from}
    var sendto,emailto,emailcc []string;
    for _,t := range to{
	tmp := mail.Address{"",t}
	emailto = append(emailto,tmp.String())
	sendto = append(sendto,tmp.Address)
    }

    for _,c := range cc{
	tmp := mail.Address{"",c}
	emailcc = append(emailcc,tmp.String())
	sendto = append(sendto,tmp.Address)
    }

    hdr := make(map[string]string)
    hdr["From"] = emailfrom.String()
    hdr["To"] = strings.Join(emailto,",")
    hdr["Cc"] = strings.Join(emailcc,",")
    // "=?UTF-8?X?  X -> B for Base64, Q for Quoted-Printable 
    hdr["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=",base64.StdEncoding.EncodeToString([]byte(subject)))
    hdr["MIME-Version"] = "1.0"
    hdr["Content-Type"] = "text/html; charset=UTF-8"
    hdr["Content-Transfer-Encoding"] = "base64"

    msg := ""
    for k,v := range hdr {
	msg += fmt.Sprintf("%s: %s\r\n",k,v)
    }
    msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(content)) 
    srv := serverip + ":" + port
    err := smtp.SendMail(srv,nil,emailfrom.Address,sendto,[]byte(msg))
    return err
}
