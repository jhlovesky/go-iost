package dpos

import (
	. "github.com/iost-official/prototype/consensus/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/iost-official/prototype/account"
	"github.com/iost-official/prototype/core/block"
)

func TestGlobalStaticProperty(t *testing.T) {
	Convey("Test of witness lists of static property", t, func() {
		prop := NewGlobalStaticProperty(account.Account{"id0", []byte{}, []byte{}}, []string{"id1", "id2", "id3"})
		Convey("New", func() {
			So(prop.NumberOfWitnesses, ShouldEqual, 3)
		})

		prop.AddPendingWitness("id4")
		prop.AddPendingWitness("id5")
		Convey("Add pending witness", func() {
			So(len(prop.PendingWitnessList), ShouldEqual, 2)
			err := prop.AddPendingWitness("id4")
			So(err, ShouldNotBeNil)
		})

		Convey("Update lists", func() {
			prop.UpdateWitnessLists([]string{"id3", "id5", "id1"})
			So(prop.WitnessList[0], ShouldEqual, "id1")
			So(prop.WitnessList[1], ShouldEqual, "id3")
			So(prop.WitnessList[2], ShouldEqual, "id5")
			So(prop.PendingWitnessList[0], ShouldEqual, "id2")
			So(prop.PendingWitnessList[1], ShouldEqual, "id4")
		})

		Convey("Delete pending witness", func() {
			err := prop.DeletePendingWitness("id4")
			So(len(prop.PendingWitnessList), ShouldEqual, 1)
			So(err, ShouldBeNil)

			err = prop.DeletePendingWitness("id2")
			So(err, ShouldNotBeNil)
		})
	})
}

func TestGlobalDynamicProperty(t *testing.T) {
	Convey("Test of global dynamic property", t, func() {
		sp := NewGlobalStaticProperty(account.Account{"id1", []byte{}, []byte{}}, []string{"id0", "id1", "id2"})
		dp := NewGlobalDynamicProperty()
		dp.LastBlockNumber = 0
		dp.TotalSlots = 0
		dp.LastBlockTime = Timestamp{0}
		startTs := Timestamp{70000}
		bh := block.BlockHead{
			Number:  1,
			Time:    startTs.Slot,
			Witness: "id0",
		}
		dp.Update(&bh)

		Convey("update first block", func() {
			So(dp.LastBlockNumber, ShouldEqual, 1)
			So(dp.TotalSlots, ShouldEqual, 1)
		})

		curSec := startTs.ToUnixSec() + 1
		sec := TimeUntilNextSchedule(&sp, &dp, curSec)
		Convey("in other's slot", func() {
			curTs := GetTimestamp(curSec)
			wit := WitnessOfTime(&sp, &dp, curTs)
			So(wit, ShouldEqual, "id0")
			So(sec, ShouldBeLessThanOrEqualTo, 3)
		})

		curSec += 2
		timestamp := GetTimestamp(curSec)
		Convey("in self's slot", func() {
			wit := WitnessOfTime(&sp, &dp, timestamp)
			So(wit, ShouldEqual, "id1")
			wit = WitnessOfSec(&sp, &dp, curSec)
			So(wit, ShouldEqual, "id1")
		})

		bh.Number = 2
		bh.Time = timestamp.Slot
		bh.Witness = "id1"
		dp.Update(&bh)
		Convey("update second block", func() {
			So(dp.LastBlockNumber, ShouldEqual, 2)
			So(dp.TotalSlots, ShouldEqual, 2)
		})

		curSec += 1
		sec = TimeUntilNextSchedule(&sp, &dp, curSec)
		Convey("in self's slot, but finished", func() {
			So(sec, ShouldBeGreaterThanOrEqualTo, 6)
			So(sec, ShouldBeLessThanOrEqualTo, 9)
		})

		curSec += 8
		Convey("in self's slot and lost two previous blocks", func() {
			curTs := GetTimestamp(curSec)
			wit := WitnessOfTime(&sp, &dp, curTs)
			So(wit, ShouldEqual, "id1")
			wit = WitnessOfSec(&sp, &dp, curSec)
			So(wit, ShouldEqual, "id1")
		})

		timestamp = GetTimestamp(curSec)
		bh.Number = 3
		bh.Time = timestamp.Slot
		bh.Witness = "id1"
		dp.Update(&bh)
		Convey("update third block", func() {
			So(dp.LastBlockNumber, ShouldEqual, 3)
			So(dp.TotalSlots, ShouldEqual, 5)
		})
	})
}
