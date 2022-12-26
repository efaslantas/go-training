# Go Ortam Kurulum ve Başlangıç Adımları


1- Kurulum : https://go.dev/dl/

2- Visual Studio Go Extension Kur
	

3- Varsayılan Path Değiştirme

    Varsayılan Go Path : %USERPROFILE%\go
    Github Reposunu üzerinde çalışmak için : go mod init github.com/efaslantas/go-training

4- İlk Kod ve Açıklamaları

    package main

import “fmt”

func main(){
    fmt.Println(“Merhaba Dünya!”)
}


package: İçerisinde package değeri aynı olan kod sayfalarının birbirleriyle iletişim kurabilirler. Bunu package sayesinde sağlar.
import: Go Kütüphanesinden örnek olarak fmt'yi koda import ettik, koda eklemek istediğiniz işlemleri import ile sağlarız
func main: Programın çalışacağı ana fonksiyondur, Derlenen uygulama ilk main fonksiyonuna bakar ve kodları {parantez} iöine yazarız

Çalıştırabilir dosya oluşturmak için : go build main.go -- Olduğunuz dizine Windowsta exe oluşturur.


