package main

import "fmt"

type SaftyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	rocksClimbed int
	sp           SaftyPlacer
}

func newRockClimber(sp SaftyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {
	//db
	//api
	//nosql
}

func (sp IceSafetyPlacer) placeSafeties() {

	fmt.Println("placing my ICE safeies...")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing my  safeies...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
