export namespace biz {
	
	export class Column {
	    prop: string;
	    label: string;
	    extractor_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Column(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.prop = source["prop"];
	        this.label = source["label"];
	        this.extractor_path = source["extractor_path"];
	    }
	}
	export class Msg {
	    is_client: boolean;
	    key_values: Record<string, string[]>;
	
	    static createFrom(source: any = {}) {
	        return new Msg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_client = source["is_client"];
	        this.key_values = source["key_values"];
	    }
	}

}

export namespace monitor {
	
	export class Filter {
	    Path: string;
	    ExpectKey: string;
	    ExpectVal: string;
	    Op: string;
	
	    static createFrom(source: any = {}) {
	        return new Filter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.ExpectKey = source["ExpectKey"];
	        this.ExpectVal = source["ExpectVal"];
	        this.Op = source["Op"];
	    }
	}
	export class Rule {
	    SrcPath: string;
	    DstKey: string;
	    DstPath: string;
	    Fs: Filter[];
	
	    static createFrom(source: any = {}) {
	        return new Rule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SrcPath = source["SrcPath"];
	        this.DstKey = source["DstKey"];
	        this.DstPath = source["DstPath"];
	        this.Fs = this.convertValues(source["Fs"], Filter);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

