package main;import b"os"
func main(){r,_:=b.ReadFile(b.Args[1])
a,i,n:=[3e4]byte{},0,0
f:u,c,k:=a[n],r[i],byte(1)
switch c{case 44:b.Stdin.Read(a[n:-^n])
case 46:print(string(u))
case 43,45:a[n]+=44-c
case 60,62:n=(n+int(c)+29939)%3e4}
for c==91&&k>0&&u<1{if i++;r[i]-c<3{k+=92-r[i]}}
for c==93&&k*u>0{if i--;r[i]-91<3{k-=92-r[i]}}
if i++;len(r)>i{goto f}}