package cli

import "fmt"

func (c *cli) TrackHandler(i interface{}) {
	ii, _ := i.([]interface{})

	if len(ii) != 3 {
		return
	}

	i2, _ := ii[2].([]interface{})
	if len(i2) < 1 {
		return
	}

	i21, _ := i2[0].([]byte)
	key := string(i21)
	if key == "" {
		return
	}

	i0, _ := ii[0].([]byte)
	if string(i0) != "message" {
		return
	}

	i1, _ := ii[1].([]byte)
	if string(i1) != "__redis__:invalidate" {
		return
	}

	fmt.Println()
	fmt.Println("refresh memory cache for key", key)
	c.use.Refresh(key)
	fmt.Println(key, "refreshed")
	fmt.Print("> ")
}
