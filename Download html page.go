package main

// import (
//     "fmt"
//     "net/http"
//     "io/ioutil"
//     "os"
//     "regexp"
//     "strconv"
//     "time"
//     )
//
// func ReturnSitesHTML(siteAdress string) string {
//   response, err := http.Get(siteAdress)
//   if err != nil {
//     fmt.Printf("%s", err)
//     os.Exit(1)
//     return ""
//   }else {
//     defer response.Body.Close()
//     contents, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//       fmt.Printf("%s", err)
//       os.Exit(1)
//       return ""
//     }
//     return string(contents)
//   }
// }
//
// func FindFoundationDate(sitesHTML string) time {
//   s := sitesHTML
//   x := regexp.MustCompile("\"post-").FindStringIndex(s)[0] + 6
//   pName := string(s[x+0]) + string(s[x+1]) + string(s[x+2]) + string(s[x+3]) + string(s[x+4]) + string(s[x+5]) + string(s[x+6]) + string(s[x+7]) + string(s[x+8])
//   response, err := http.Get(string("http://vk.com/wall-" + pName + "_1")) //Первая могла быть удалена.
//   if err != nil {
//     fmt.Printf("%s", err)
//     os.Exit(1)
//     return ""
//   } else {
//     contents, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//       fmt.Printf("%s", err)
//       os.Exit(1)
//       return ""
//     }
//     return FormDateInTimeType((regexp.MustCompile("[0-9]{1,2} [а-я]{3} [0-9]{4}").FindString(string(contents))))
//   }
// }
//
// /*func main () {
//   siteAdress := "http://vk.com/p2pworld"
//   file, err := os.Create("test.txt")
//   if err != nil {
//     // handle the error here
//     return
//   }
//   defer file.Close()
//   file.WriteString((ReturnSitesHTML(siteAdress)))
//   fmt.Println(FindFoundationDate(ReturnSitesHTML(siteAdress)))
// }*/
//
// func FormDateInTimeType (date string) time {
//   return time.Date(FormYear(date), FormMonth(date), FormDay(date), 0, 0, 0, 0, time.UTC)
// }
//
// func FormYear(date string)  int{
//   return strconv.Itoa(regexp.MustCompile("[0-9]{4}").FindString(date))
// }
//
// func FormDay(date string)  int{
//   return strconv.Itoa(regexp.MustCompile("[0-9]{1,2}").FindString(date))
// }
//
// func FormMonth(date string)  int{
//   s := regexp.MustCompile("[а-я]{3}").FindString(date)
//   switch s {
//   case "янв":
//     return time.January
//   case "фев":
//     return time.February
//   case "мар":
//     return time.March
//   case "апр":
//     return time.April
//   case "май":
//     return time.May
//   case "июн":
//     return time.June
//   case "июл":
//     return time.July
//   case "авг":
//     return time.August
//   case "сеп":
//     return time.September
//   case "окт":
//     return time.October
//   case "ноя":
//     return time.November
//   default :
//     return time.December
//   }
// }
