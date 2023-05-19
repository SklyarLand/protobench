package main

import (
	"encoding/json"
	"fmt"
	"github.com/francoispqt/gojay"
	"github.com/mailru/easyjson"
	"testing"

	"github.com/golang/protobuf/proto"
	"ptotobench/pb"
)

var (
	PBSmall = pb.BenchSmall{
		Action: "benchmark",
		Key:    "data to be sent",
	}

	PBMedium = pb.BenchMedium{
		Name:   "Tester",
		Age:    20,
		Height: 5.8,
		Weight: 180.7,
		Alive:  true,
		Desc: `If you’ve ever heard of ProtoBuf you may be thinking 
		that the results of this benchmarking experiment will be obvious,
		JSON < ProtoBuf.`,
	}

	PBLarge = pb.BenchLarge{
		Name:     "Tester",
		Age:      20,
		Height:   5.8,
		Weight:   180.7,
		Alive:    true,
		Desc:     "Lets benchmark some json and protobuf",
		Nickname: "Another name",
		Num:      2314,
		Flt:      123451231.1234,
		Data: `If you’ve ever heard of ProtoBuf you may be thinking that
		the results of this benchmarking experiment will be obvious, JSON < ProtoBuf.
		My interest was in how much they actually differ in practice.
		How do they compare on a couple of different metrics, specifically serialization
		and de-serialization speeds, and the memory footprint of encoding the data.
		I was also curious about how the different serialization methods would
		behave under small, medium, and large chunks of data.`,
	}
)

func TestMarshal(_ *testing.T) {
	fmt.Printf("Marshal Gojay ---------\n")
	bs := PBLarge
	g, _ := gojay.MarshalJSONObject(&bs)
	fmt.Println(string(g))

	fmt.Printf("Marshal Json ---------\n")
	g, _ = json.Marshal(&bs)
	fmt.Println(string(g))

	fmt.Printf("Marshal Easyjson ---------\n")
	g, _ = easyjson.Marshal(&bs)
	fmt.Println(string(g))

	fmt.Printf("\n")
}

func TestDataAllocationsSmall(_ *testing.T) {
	fmt.Printf("Small ---------\n")
	bs := PBSmall
	j, _ := json.Marshal(&bs)
	p, _ := proto.Marshal(&bs)
	g, _ := gojay.Marshal(&bs)
	e, _ := easyjson.Marshal(&bs)

	printInfo(j, "json")
	printInfo(g, "gojay")
	printInfo(e, "easyjson")
	printInfo(p, "protobuf")
	fmt.Printf("\n")
}

func TestDataAllocations(_ *testing.T) {
	fmt.Printf("Medium ---------\n")
	bs := PBMedium
	j, _ := json.Marshal(&bs)
	p, _ := proto.Marshal(&bs)
	g, _ := gojay.Marshal(&bs)
	e, _ := easyjson.Marshal(&bs)

	printInfo(j, "json")
	printInfo(g, "gojay")
	printInfo(e, "easyjson")
	printInfo(p, "protobuf")
	fmt.Printf("\n")
}

func TestDataAllocationsLarge(_ *testing.T) {
	fmt.Printf("Large ---------\n")
	bs := PBLarge
	j, _ := json.Marshal(&bs)
	p, _ := proto.Marshal(&bs)
	g, _ := gojay.Marshal(&bs)
	e, _ := easyjson.Marshal(&bs)

	printInfo(j, "json")
	printInfo(g, "gojay")
	printInfo(e, "easyjson")
	printInfo(p, "protobuf")
	fmt.Printf("\n")
}

func BenchmarkJSONMarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&s)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&m)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&l)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkGojayMarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := gojay.Marshal(&s)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := gojay.Marshal(&m)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := gojay.Marshal(&l)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkEasyJsonMarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := easyjson.Marshal(&s)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := easyjson.Marshal(&m)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := easyjson.Marshal(&l)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufMarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&s)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&m)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&l)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	sd, _ := json.Marshal(&s)
	md, _ := json.Marshal(&m)
	ld, _ := json.Marshal(&l)

	var sf pb.BenchSmall
	var mf pb.BenchMedium
	var lf pb.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(sd, &sf)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(md, &mf)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(ld, &lf)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkGojayUnmarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	sd, _ := gojay.Marshal(&s)
	md, _ := gojay.Marshal(&m)
	ld, _ := gojay.Marshal(&l)

	var sf pb.BenchSmall
	var mf pb.BenchMedium
	var lf pb.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = gojay.Unmarshal(sd, &sf)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = gojay.Unmarshal(md, &mf)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = gojay.Unmarshal(ld, &lf)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkEasyJsonUnmarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	sd, _ := easyjson.Marshal(&s)
	md, _ := easyjson.Marshal(&m)
	ld, _ := easyjson.Marshal(&l)

	var sf pb.BenchSmall
	var mf pb.BenchMedium
	var lf pb.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = easyjson.Unmarshal(sd, &sf)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = easyjson.Unmarshal(md, &mf)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = easyjson.Unmarshal(ld, &lf)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	s := PBSmall
	m := PBMedium
	l := PBLarge

	sd, _ := proto.Marshal(&s)
	md, _ := proto.Marshal(&m)
	ld, _ := proto.Marshal(&l)

	var sf pb.BenchSmall
	var mf pb.BenchMedium
	var lf pb.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(sd, &sf)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(md, &mf)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(ld, &lf)
		}
	})
}

func printInfo(d []byte, ser string) {
	used := len(d)
	allocated := cap(d)
	fmt.Printf("Type: %s \t\tData size: %d \t\tTotal Allocated: %d \t\t Used/Allocated: %.2f%%\n", ser, used, allocated, percentUsed(used, allocated)*100)
}

func percentUsed(used, allocated int) float32 {
	return float32(used) / float32(allocated)
}
