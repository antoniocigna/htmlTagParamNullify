package main


import (
	"fmt"
    "os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


var inpHTML = "inpHTML.html.txt"
var outFile = "outHTML.txt"
var outFile2 = "outText.txt"

//---------------
func nullifyTag(str1 string,tagName1 string) string {
	var len1 = len(str1); 
	tag0 := "<" + strings.TrimSpace(tagName1);
	tag1 := tag0+">"; 
	tag0 += " "; 

	var j2=0;
	var j3=0;
	
	str1 = strings.ReplaceAll(str1, strings.ToUpper(tag0), tag0); 
	
	strOut:=""
	for r:=0; r < len(str1);r++ {
		j2 = strings.Index(str1,tag0); 
		if (j2 < 0) {
			strOut+= str1;
			break;
		} 
		str2:=str1[j2:];  
		strOut += str1[:j2] + tag1; 
		j3 = strings.Index(str2,">");  
		if (j3 > 0) {
			str1 = str2[j3+1:]; 
		} else {
			str1 = str2; 
		}	
	}
	var len2 = len(strOut); 	
	fmt.Println("	trasforma " + (tag1 + "                  ")[0:10]+ " \t length: \t", len1, "\t=>\t", len2);	
	return strOut; 
}
//-----------------------
//-----------------------
func trasforma1(str1 string) string {
	str1 = nullifyTag(str1, "body");
	str1 = nullifyTag(str1, "table");
	str1 = nullifyTag(str1, "tbody");
	str1 = nullifyTag(str1, "tr");
	str1 = nullifyTag(str1, "td");	
	str1 = nullifyTag(str1, "span");
	str1 = nullifyTag(str1, "ins");
	str1 = nullifyTag(str1, "script");
	str1 = nullifyTag(str1, "div");
	str1 = strings.ReplaceAll(str1, "<span>","");
	str1 = strings.ReplaceAll(str1, "</span>","");
	str1 = strings.ReplaceAll(str1, "(adsbygoogle=window.adsbygoogle||[]).push({});","");	
	return str1; 
}
//----------------------------------
func extractTextTD(str1 string){
	var strText =""
    var tdRighe = strings.Split(str1,"<td>");
	for g:=0; g < len(tdRighe); g++ {
		tdRiga := tdRighe[g];
		j1:= strings.Index(tdRiga, "</td>"); 
		if j1 < 0 {continue}
		tdRiga = strings.TrimSpace( tdRiga[0:j1]);
		tdRiga = strings.ReplaceAll( tdRiga, "</", "<");		
		tdRiga = strings.ReplaceAll( tdRiga, "<script>","");		
		tdRiga = strings.ReplaceAll( tdRiga, "<ins>"," ");
		tdRiga = strings.ReplaceAll( tdRiga, "<div>"," ");
		tdRiga = strings.ReplaceAll( tdRiga, "&nbsp;"," ");
		tdRiga = strings.ReplaceAll( tdRiga, "<br>"," ");
		tdRiga = strings.ReplaceAll( tdRiga, "<strong>","");
		tdRiga = strings.ReplaceAll( tdRiga, "<p>"," ");
		tdRiga = strings.ReplaceAll( tdRiga, "<em>","");
		tdRiga = strings.TrimSpace( tdRiga);
		if len(tdRiga) < 2  {continue;}
		
		tdRiga += ".";
		
		tdRiga = strings.ReplaceAll( tdRiga, "..",".");
		tdRiga = strings.ReplaceAll( tdRiga, "?.","?");		
		tdRiga = strings.ReplaceAll( tdRiga, "!.","!");
	
		strText += tdRiga + "\n";
	}
	f, err := os.Create(outFile2)
    check(err)
    defer f.Close()
    _, err = f.WriteString( strText )
    check(err)
	f.Sync();
	
} // end of extractTextTD
//---------------------------------------------
func main() {
	dat, err := os.ReadFile( inpHTML)
    check(err)
	
	str1 := string(dat);
	
	fmt.Println("legge stringa lunga " , len(str1) ); 
	
	str1 = trasforma1( str1 );
	
	extractTextTD(str1); 
	
	fmt.Println("scrive stringa lunga " , len(str1) ); 
	
	f, err := os.Create(outFile)
    check(err)
    defer f.Close()
    _, err = f.WriteString( str1 )
    check(err)
	f.Sync()	
	
	
} // end of main 
//----------------------------------------------