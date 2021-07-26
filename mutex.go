package main

import ("fmt"
  "sync"
)



var (
  mutex sync.Mutex
  money int 
)

func init(){
  money = 2500
}

func John(value int, wg *sync.WaitGroup){
  mutex.Lock()
  fmt.Printf("John is giving me %d :%d\n", value, money)
  money += value
  mutex.Unlock()
  wg.Done()
}

func Samuel(value int, wg *sync.WaitGroup ){
  mutex.Lock()
  fmt.Printf("Samuel is taking from me %d :%d\n", value, money)
  money -= value
  mutex.Unlock()
  wg.Done()
}

func main(){
  fmt.Println("start mutext")

  var wg sync.WaitGroup

  wg.Add(2)
  go John(500, &wg)
  go Samuel(800, &wg)
  wg.Wait()

  fmt.Printf("Total money now %d", money)

}