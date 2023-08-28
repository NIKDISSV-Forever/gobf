package B
import "os"
func R(a[]byte){for i,n,f:=0,0,[3e4]byte{};i<len(a);i++{u,c,k:=f[n],a[i],byte(1)
switch c{case 44:os.Stdin.Read(f[n:-^n])
case 46:print(string(u))
case 43,45:f[n]+=44-c
case 60,62:n=(n+int(c)+29939)%3e4}
for c==91&&k>0&&u<1{if i++;a[i]-c<3{k+=92-a[i]}}
for c==93&&k*u>0{if i--;a[i]-91<3{k-=92-a[i]}}}}