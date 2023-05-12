package struct2ts_test

import (
	"encoding/json"
	"os"
	"time"

	"github.com/OneOfOne/struct2ts"
)

type OtherStruct struct {
	T time.Time `json:"t,omitempty"`
}

type ComplexStruct[K ~float32 | ~float64] struct {
	S           string           `json:"s,omitempty"`
	I           int              `json:"i,omitempty"`
	F           float64          `json:"f,omitempty"`
	TS          *int64           `json:"ts,omitempty" ts:"date,null"`
	T           time.Time        `json:"t,omitempty"` // automatically handled
	NullOther   *OtherStruct     `json:"o,omitempty"`
	NoNullOther *OtherStruct     `json:"nno,omitempty" ts:",no-null"`
	Data        Data             `json:"d"`
	DataPtr     *Data            `json:"dp"`
	RawPtr      *json.RawMessage `json:"rm"`
	GenericTest K                `json:"genericTest"`
}

type Float32Alias float32

type Data map[string]interface{}

func ExampleComplexStruct() {
	s2ts := struct2ts.New(nil)
	s2ts.Add(ComplexStruct[Float32Alias]{})
	s2ts.RenderTo(os.Stdout)

	// Output:
	// // helpers
	// const maxUnixTSInSeconds = 9999999999;
	//
	// function ParseDate(d: Date | number | string): Date {
	// 	if (d instanceof Date) return d;
	// 	if (typeof d === 'number') {
	// 		if (d > maxUnixTSInSeconds) return new Date(d);
	// 		return new Date(d * 1000); // go ts
	// 	}
	// 	return new Date(d);
	// }
	//
	// function ParseNumber(v: number | string, isInt = false): number {
	// 	if (!v) return 0;
	// 	if (typeof v === 'number') return v;
	// 	return (isInt ? parseInt(v) : parseFloat(v)) || 0;
	// }
	//
	// function FromArray<T>(Ctor: { new (v: any): T }, data?: any[] | any, def = null): T[] | null {
	// 	if (!data || !Object.keys(data).length) return def;
	// 	const d = Array.isArray(data) ? data : [data];
	// 	return d.map((v: any) => new Ctor(v));
	// }
	//
	// function ToObject(o: any, typeOrCfg: any = {}, child = false): any {
	// 	if (o == null) return null;
	// 	if (typeof o.toObject === 'function' && child) return o.toObject();
	//
	// 	switch (typeof o) {
	// 		case 'string':
	// 			return typeOrCfg === 'number' ? ParseNumber(o) : o;
	// 		case 'boolean':
	// 		case 'number':
	// 			return o;
	// 	}
	//
	// 	if (o instanceof Date) {
	// 		return typeOrCfg === 'string' ? o.toISOString() : Math.floor(o.getTime() / 1000);
	// 	}
	//
	// 	if (Array.isArray(o)) return o.map((v: any) => ToObject(v, typeOrCfg, true));
	//
	// 	const d: any = {};
	//
	// 	for (const k of Object.keys(o)) {
	// 		const v: any = o[k];
	// 		if (v === undefined) continue;
	// 		if (v === null) continue;
	// 		d[k] = ToObject(v, typeOrCfg[k] || {}, true);
	// 	}
	//
	// 	return d;
	// }
	//
	// // structs
	// // struct2ts:github.com/OneOfOne/struct2ts_test.OtherStruct
	// class OtherStruct {
	// 	t: Date;
	//
	// 	constructor(data?: any) {
	// 		const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
	// 		this.t = ('t' in d) ? ParseDate(d.t) : new Date();
	// 	}
	//
	// 	toObject(): any {
	// 		const cfg: any = {};
	// 		cfg.t = 'string';
	// 		return ToObject(this, cfg);
	// 	}
	// }
	//
	// // struct2ts:github.com/OneOfOne/struct2ts_test.ComplexStruct
	// class ComplexStruct {
	// 	s: string;
	// 	i: number;
	// 	f: number;
	// 	ts: Date | null;
	// 	t: Date;
	// 	o: OtherStruct | null;
	// 	nno: OtherStruct;
	// 	d: { [key: string]: any };
	// 	dp: { [key: string]: any } | null;
	// 	rm: any;
	// 	genericTest: number;
	//
	// 	constructor(data?: any) {
	// 		const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
	// 		this.s = ('s' in d) ? d.s as string : '';
	// 		this.i = ('i' in d) ? d.i as number : 0;
	// 		this.f = ('f' in d) ? d.f as number : 0;
	// 		this.ts = ('ts' in d) ? ParseDate(d.ts) : null;
	// 		this.t = ('t' in d) ? ParseDate(d.t) : new Date();
	// 		this.o = ('o' in d) ? new OtherStruct(d.o) : null;
	// 		this.nno = new OtherStruct(d.nno);
	// 		this.d = ('d' in d) ? d.d as { [key: string]: any } : {};
	// 		this.dp = ('dp' in d) ? d.dp as { [key: string]: any } : null;
	// 		this.rm = ('rm' in d) ? d.rm as any : null;
	// 		this.genericTest = ('genericTest' in d) ? d.genericTest as number : 0;
	// 	}
	//
	// 	toObject(): any {
	// 		const cfg: any = {};
	// 		cfg.i = 'number';
	// 		cfg.f = 'number';
	// 		cfg.t = 'string';
	// 		cfg.genericTest = 'number';
	// 		return ToObject(this, cfg);
	// 	}
	// }
	//
	// // exports
	// export {
	// 	OtherStruct,
	// 	ComplexStruct,
	// 	ParseDate,
	// 	ParseNumber,
	// 	FromArray,
	// 	ToObject,
	// };
}
