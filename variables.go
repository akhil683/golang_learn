// bool, string, int, float64, byte

package main
import "fmt"

func main() {
  // Initialize variables
  var smsSendingLimit int
  var costPerSMS float64
  var hasPermission bool
  var username string

  // Initialize GOATed variables
  // := walrus operator
  mySkillIssues := 42

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username, mySkillIssues)
}


