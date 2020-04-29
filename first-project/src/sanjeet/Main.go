/*
$gedit ~/.bashrc

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

export GOPATH=/home/sanjeet/golib
export PATH=$PATH:$GOPATH/bin
export GOPATH=$GOPATH:/home/sanjeet/code
$source /.bashrc
//download threee main lib for go
$go get -u github.com/nsf/gocode
$go get -u github.com/rogpeppe/godef
$go get -u golang.org/x/tools/cmd/guru


//eclipse configure
go install : /usr/local/go
gopath :  :/home/sanjeet/golib
gofmt : /usr/local/go/bin/gofmt
gocode : /home/sanjeet/golib/bin/gocode
godef : /home/sanjeet/golib/bin/godef
guru : /home/sanjeet/golib/bin/gocode
//to update gocode,godef,gocode
git pull --ff-only
// running golib project
$go build package_name
$package_name


*/ 

package main
import "fmt";
//  import "reflect"
//import "net/http"
import "log"
import "bytes"
//import "io"
import "sync"
import "time"
//import "runtime"

const(
	ca=iota
	cb=iota
	cc

)
const(
	_=iota//ignore first case
	KB=1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB

)
const(
	isAdmin=1<<iota
	isleader
	isdeveloper
	
	Africa
	Asia
)


type Doctor struct{
	number int
	name string
	companions []string

}

type Animal struct{
	Name string
	Origin string

}
type Bird struct{
	Animal
	Feather float32
	Fly bool

}
type Car struct{
Name string `required max:"100"`
Origin string

}
/* goroutes */

var wg=sync.WaitGroup{}
var counter=0
var m= sync.RWMutex{}//readwritemutex
/* channels */

const (
	logInfo="info"
	logWarn ="warning"
	logErr ="error"

)
type logEntry struct{

	time time.Time
	severity string
	message string
}

var logch=make(chan logEntry,45)
var donech=make(chan struct{})


func main(){
	/*var i int;
	i = 42
	var j float32 =27;
	k := 99.00
	n:=1==1;
	var p bool;
	var u uint16=44;
	fmt.Printf("%v,%T \n",i,i);
	fmt.Printf("%v ,%T \n",j,j);
	fmt.Printf("%v,%T \n",k,k);
	fmt.Printf("%v ,%T\n",n,n);
	fmt.Printf("%v ,%T\n",p,p);
	fmt.Printf("%v ,%T\n",u,u);*/
	/*
	a:=10
	b:=3
	fmt.Println(a&b)
	fmt.Println(a|b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b) */
	
	/* bit shifting
	a:=9
	fmt.Println(a<<3)//2^3*2^3=2^6
	fmt.Println(a>>3)//2^3 /2^3 =2^1
	
	*/
	/*
	n:=3.14
	n=13.7e72
	n=2.1E14
	fmt.Printf("%v,%T",n,n)*/
	
	/*
	var n complex64=1+2i
	fmt.Printf("%v,%T \n",n,n)
	a:=1+2i
	b:=2+5.2i
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	fmt.Println(a + b)
	var p complex128=complex(5,25)
	fmt.Printf("%v,%T",p,p)*/
	/*
	s:= "this is sanjeet "
	s2:="pal"
	fmt.Printf("%v,%T\n",string(s[2]),s[2])
	fmt.Println(s+s2)
	b:=[]byte(s)
	fmt.Printf("%v,%T\n",b,b)
	var r rune='a'
	fmt.Printf("%v,%T",r,r)*/
	
	
	/* constant */
	/*const pal int = 42
	fmt.Println(pal)
	const a=42
	fmt.Printf("%v,%T\n",a,a)
	//fmt.Prinf("%v,%T",pal+a)//error
	var b int16=27
	fmt.Printf("%v,%T\n",a+b,a+b)
	*/
	/* working with iota*/
	/*fmt.Printf("%v,%T,%v,%v\n",ca,ca,cb,cc)
	
	fileSize:=400000000000.
	fmt.Printf("%.2fGB\n",fileSize/GB)
	
	var roles byte= isAdmin | Africa | Asia
	fmt.Printf("%b,%v\n",roles,roles)
	
	fmt.Printf("isAdmin..%v\n",isAdmin & roles == isAdmin)
	fmt.Printf("is leader..%v",isleader& roles==isleader)
	*/
	
	/* Array */
	/* students :=[45]string{"sanjeet","pal"}
	 students_grades:=[...]int{78,45,12}
	 var my_friend [3]string
	 my_friend[0]="rahul"
	 my_friend[1]="mohan"
	 my_friend[2]="sohan"
	 fmt.Printf("Student: %v\n",students[0])
	 fmt.Printf("student grades%v\n",students_grades);
	 fmt.Printf("my_friends %v ,%v\n",my_friend[0],len(my_friend))
	 var identimetrix [3][3]int=[3][3]int{[3]int{2,1,5},[3]int{4,5,6},[3]int{7,8,9}}
	 fmt.Println(identimetrix)
	 
	 copy_identimatrix:=&identimetrix
	 fmt.Println(copy_identimatrix)
	  //slices
	  palx := []int{45,78,98,4,7815,45,45,15,21,5,1}
	  palx=append(palx,7845,454124)
	  palx=append(palx,[]int{2,3,4,5}...)
	  palx=palx[1:]//pop operaiton
	  palx=palx[:len(palx)-1] //pop from last
	  
	  palx=append(palx[:3],palx[4:]...)//removing 7815 loc(4) cannot use with copy
	  
	  copy_palx := palx
	  palx_a:=palx[:]//slice all element
	  palx_b:=palx[3:]//lice element from 3rd element to end
	  palx_c:=palx[:6]//first 6 element
	  palx_d:=palx[3:6]//slice 4 ,5,6 element
	  fmt.Println(pal)
	  fmt.Printf("capacity:%v",cap(palx))
	  fmt.Println(copy_palx)
	  fmt.Println(palx_a)
	  fmt.Println(palx_b)
	  fmt.Println(palx_c)
	  fmt.Println(palx_d)
	
	 make_func := make([]int,3,100) //100 is capacity,3 i length //capacity inc by double
	 fmt.Println(make_func)
	*/
	
	/* maps and structs*/
	/*
	statePopulation := make(map[string]int,10)
	statePopulation = map[string]int{
		"california": 39250017,
		"Texas": 2789456,
		"Florida":454512,
		"new york":4541541,
		"pennysylvania":452421,
		"Illinois":45415,
		"ohio":4545,
		
	
	}
	statePopulation["india"]=4512121210213121
	fmt.Println(statePopulation)
	delete(statePopulation,"new york")
	fmt.Println(statePopulation["ohio"])
	fmt.Println(statePopulation)
	pop,ok:=statePopulation["india"]
	fmt.Println(pop,ok) 
	
	doctor_object:=Doctor{
		number:3,
		name:"sanjeet pal",
		companions: []string{
			"pal.org",
			"pal.com",
			"pal.in",
		},
	}
	fmt.Println(doctor_object.companions[1])
	 
	oStudent:=struct{name string}{name:"sanjeet pal"}
	anotherStudent:=oStudent
	anotherStudent.name ="saurab sparrow"
	fmt.Println(oStudent)
	fmt.Println(anotherStudent)
	
	
	//inheritance
	bird:=Bird{}
	bird.Name="penguin"
	bird.Origin="antartica"
	bird.Feather=2
	bird.Fly=false
	fmt.Println(bird)
	
	bird:=Bird{
		Animal:Animal{Name:"penguin",Origin:"Antartica"},
		Feather:2,
		Fly:false,
	}
	fmt.Println(bird.Name)
	
	//validation 
	car:=reflect.TypeOf(Car{})
	field, _:=car.FieldByName("Name")
	fmt.Println(field.Tag)

	*/
	
	
	/* CONTROL FLOW */
	/*if true{
		fmt.Println("This is if statement..")
	}
	statePopulation:=map[string]int{
	"california":85415412,
	"Texas":4545464,
	"Ohio":515212,
	"New York":5215241,
	"Pennsylvania":54461212,
	
	}
	if pop,ok :=statePopulation["Ohio"];ok{
		fmt.Println(pop)
		
	}
	fmt.Println(statePopulation)
	number :=50
	guess :=-45
	if guess >=1 && guess <=100{
		if guess<number{
			fmt.Println("low")
		}
		if guess >number{
			fmt.Println("high")
		}
		if guess==number{ 
			fmt.Println("equals")
		}
	
	}
	fmt.Println(defaultfunc())
	switch  i:=1+1;i{
		case 1,5:
			fmt.Println("one")
			fallthrough //logic pass to next case
			
		case 2,56,6:
			fmt.Println("two")
			break //logic dodnot pass to next case
		default:
			fmt.Println("default")		

	}
	
	var i interface{} =[3]int{}
	switch i.(type){
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is  a float64")
	case [2]int:
		fmt.Println("this is int")
	case string:
		fmt.Println("i is  a string")
	default:
		fmt.Println("i is another type")			
	
	}
	*/
	 /* Lopping */
	 /*
	
	for i,j:=0,0;i<5;i,j=i+1,j+2{
		fmt.Println(i,j)
	}
	//while loop
	
	i:=0
	for i<5{
		fmt.Println(i)
		i++
	}
	for i:=0;i<10;i++{
		if i%2==0{
			continue
		}
		fmt.Print(i)
	}
	Loop:
		for i:=1;i<=100;i++{
			for j :=1;j<=50;j++{
				fmt.Print(j)
				if(j==12){
					break Loop
				}
			}
		}
	
	s:=[]int{10,45,78}
	for k,v := range s{ 
		//this forloop is work for map,string,arrays,slices,channel
		fmt.Println(k,v)
	}
	*/
	/*Defer ,Panic ,Recover */
	
	/*fmt.Println("start")
	defer fmt.Println("middle") //it will be executed after every statement is executed
	fmt.Println("end")
	*/
	//hello world go program
	/*http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("hello Go!"))
	
	})
	err := http.ListenAndServe(":8880",nil)
	if(err!=nil){
	panic(err.Error())
	}*/
	
	/*fmt.Println("start")
	defer fmt.Println("this is deferred")
	panic("This is panic statement")//executed last after defer and no futher statement will be executed
	fmt.Println("end")
	*/
	
	/*fmt.Println("start")
	panicker()
	fmt.Println("end")*/
	
	/* Pointer */
	//creating
	//Deferencing
	//working nil
	//internal Pointer
/*	
	var x int =42
	var y *int=&x
	fmt.Println(&x,y,*y,x)
	*y=77
	fmt.Println(&x,y,*y,x)
	
	a:=[3]int{1,2,3}
	b:=&a[0]
	c:=&a[1]
	fmt.Println(a,b,c,&a,&b,&c)
	
	var ms *pointer_struct
	fmt.Println(ms)
	ms=new(pointer_struct)
	//(*ms).foo=78 //reference 
	//fmt.Println((*ms).foo)
	ms.foo=45
	fmt.Println(ms.foo)
	//map and slices are pointer that point to underlying data but doesn't contain the data itself 
	//so making changes will be reflected in pojected variable
	*/
	
	/* functions */
	/*roll_no:=89
	
	message("hello world",&roll_no)
	sx:=sum(1,2,3,4,5,6)
	fmt.Println(sx)
	
	//errors
	/*dx,err := divide(9.6,0.0)
	if	err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(dx)
	//anonymous function
	func(){
		fmt.Println("anonymous function")
	}() //()resolve problem function never used   
	
	//fun under for loop
	for i:=0;i<5;i++{
		func(i int){
			fmt.Println(i)
		}(i)
	}
	var f func()=func(){
		fmt.Println("Hello Go")
	}
	f()
	/*
	var multiply func(float64,float64)(float64,error)
	multiply = func(a,b float64)(float64,error){
		if(b==0.0){
		return 0.0,fmt.Errorf("Cannot divide by zero")
		}
		return a/b ,nil
	}
	
	
	d,err := multiply(6.2,0.0)
	if err!= nil{
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	
	// struct and function
	g := greeter{
		greeting : "snajeet",
		name :"pal",
	}
	g.greet()
	*/
	 /* Interfaces */
	  //{io.Writer,io.Reader,interface{}}
	//composing conversion
	//type conversion
	//--empty interface
	//--type switches
	//implementing with values vs pointer
	/*var f Flower = ConsoleFlower{}
	f.Rose([]byte("rose"))
	
	myInt := IntCounter(0)
	var inc Incrementer =&myInt
	for i:=0;i<10;i++{
		fmt.Println(inc.Increment())
	}*/
	
	/*
	var wc ReaderCloser=NewBufferedReaderCloser()
	wc.read([]byte("hello sanjeet pal,ad"))
	wc.close()
	
	//type conversion
	bwc := wc.(*BufferedReaderCloser)
	fmt.Println(bwc)
	 r,ok :=wc.(*BufferedReaderCloser)
	 if ok{
		 fmt.Println(r)
	 }else{
		 fmt.Println("Conversion failed")
	 }
	//empty interface
	var myObj interface{} =NewBufferedReaderCloser()
	if wc, ok := myObj.(ReaderCloser);ok{
		wc.read([]byte("hello sanjeet pal"))
		wc.close()
	}
	rk,ok := myObj.(io.Reader)// readmethod has  pointer receiver
	if ok{
		fmt.Println(rk)
	}else{
		fmt.Println("Conversion failed")
	}
	//simple empty interface
	
	var i interface{}=0
	switch i.(type){
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("other")	
	
	}
	//method set of value is all methods with value receivers
	//Method set of pointer is all methods,regardlessof receiver types
	*/
	
	
	/* goroutines */
	//concurrent programming
	//creating goroutines
	//synchronization
	//waitGroups
	//mutexes
	//parallelism
	
	/*
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	
	
	var message ="hello sanjeet"
	wg.Add(1)
	go func(message string){
		fmt.Println(message)
		wg.Done()
	}(message)
	message="hello pal"
	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
	*/
	//mutex is lock just like thread ,object,class lock in java
	/*fmt.Printf("number of processor/threads %v \n",runtime.GOMAXPROCS(-1))//no of thread
	for i :=0;i<10;i++{
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	//check race condition
	//go run -race src/main.go
	*/
	
	
	/* channels */
	//pass data between goroutine with safely env
	//Restricting data flow
	//BufferedChannels
	//closing channels
	//for loop channel
	
	ch :=make(chan int,50) //creating channel second param is buffered to send multiple message
 	wg.Add(2)
	go func(ch <- chan int){//receive only by using parameter ch <- chan int
		//i:= <- ch
		//fmt.Println(i)
		//i = <- ch
		//fmt.Println(i)
		
		
		//for i := range ch{
			for{
			if  i,ok:= <- ch; ok{
				fmt.Println(i)
			}else{
				break
			}
			
		}

		//ch <-89 //send
		wg.Done()
	}(ch)
	go func(ch chan <- int){//send only by using parameter ch chan <- int
		ch <- 42
		ch <-89
		close(ch)
		//fmt.Println(<-ch)//receive
		wg.Done()
	}(ch)
	wg.Wait()
	
	
	//closing channel using select 
	go logg()
	logch <- logEntry{time.Now(),logInfo,"app starting"}
	logch<- logEntry{time.Now(),logInfo,"app shutting"}
	time.Sleep(10 *time.Millisecond)
	
	
	
	
	
	
}
/* channels */
func logg(){
	for{
		select{
		case entry := <-logch:
			fmt.Println(entry.time.Format("2006-01-02T15:04:06"),entry.severity,entry.message)
		case <- donech:
			break	
		}
	
	}


}

/*goroutines */

func sayHello(){
	//m.RLock()
	fmt.Println(counter)
	fmt.Println("hello world")
	m.RUnlock()
	wg.Done()
	
}
func increment(){
	//m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}




/* interfaces */
type Reader interface{
	read([] byte)(int,error)
}
type Closer interface{
	close() error
}
type ReaderCloser interface{
	Reader
	Closer

}
type BufferedReaderCloser struct{
	buffer *bytes.Buffer
}
func (bwc *BufferedReaderCloser) read(data []byte)(int,error){
	n,err := bwc.buffer.Write(data)
	if err!=nil{
		return 0,err
	}
	v :=make([]byte,8)
	for bwc.buffer.Len() >8{
		_,err := bwc.buffer.Read(v)
		if err != nil{
			return 0,err
		}
		_,err = fmt.Println(string(v))
		if err !=nil{
			return 0,err
		}
	}
	return n,nil
}
func (bwc *BufferedReaderCloser) close() error{
	for bwc.buffer.Len()>0{
		data := bwc.buffer.Next(8)
		_,err :=fmt.Println(string(data))
		if err != nil{
			return err
		}
	}
	return nil

}

func NewBufferedReaderCloser() *BufferedReaderCloser {
	return &BufferedReaderCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
//----

type Incrementer interface{
	Increment() int
}
type IntCounter int
func (ic *IntCounter) Increment() int{
	*ic++
	return int(*ic)
}

//---







type Flower  interface{
	Rose([]byte)(int,error)
}
type ConsoleFlower struct{
}
func (cw ConsoleFlower) Rose(data []byte)(int,error){
	n,err := fmt.Println(string(data))
	return n,err
}

/* function*/
type greeter struct{
	greeting string
	name string
}
func (g greeter)greet(){
	fmt.Println(g.greeting,g.name)
}



func divide(a,b float64)(float64,error){
	if(b==0.0){
		return 0.0,fmt.Errorf("Cannot divide by zero")
	}
	return a/b ,nil

}

func sum(values ...int) int{
	fmt.Println(values)
	result:=0
	for _,v :=range values{
	
		result += v
	}
	fmt.Println("The sum is ",result)
	return result
	
	
	
}
/* pointer*/
func message(mess string,roll_no *int){
	fmt.Println(mess,*roll_no)
}
type pointer_struct struct{
foo int

}

func panicker(){
	fmt.Println("about to panic")
	defer func(){
		if err := recover();err!=nil{
			log.Println("Error",err)//Recover Only useful in deferred function
			///panic(err) //main function will get in panic
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")

}
func defaultfunc() bool{
		fmt.Println("this is default func body")
		return true
}

	
	
	
	
	
	