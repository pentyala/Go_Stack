package main

import "fmt"

var st [100]int
var top int
func main(){
    top=-1
    ch := 0
    temp:=0
    for true{
        fmt.Println("Enter you choice:")
        fmt.Println("1. PUSH\n2. POP\n3. PRINT\n4. EXIT")
        fmt.Scanf("%d\n",&ch)
        switch ch {
            case 1:
            fmt.Println("Enter the value...")
            fmt.Scanf("%d\n",&temp)
            push(temp);
            case 2:
            temp=pop()
            if temp!=-1{
                fmt.Println("The popped value is ",temp)
            }
            case 3:
            print()
            case 4:
            return
            default:
            fmt.Println("Please enter a valid choice")
        }
    }
}

func print(){
    i:=0
    if top==-1 {
        fmt.Println("First insert elements into the stack")
    }else{
        fmt.Printf("The values are as follows")
        for i<=top{
            fmt.Println(st[i])
            i++
        }
    }
}

func pop() int{
    if top==-1{
        fmt.Println("Please push values before popping")
        return -1;
    }
    temp:=st[top]
    top--
    return temp
}

func push(n int){
    top++
    st[top]=n
}