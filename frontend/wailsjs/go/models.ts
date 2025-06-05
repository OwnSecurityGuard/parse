export namespace main {
	
	export class Tmp {
	    Timestamp: number;
	    Uri: string;
	    Msg: string;
	
	    static createFrom(source: any = {}) {
	        return new Tmp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Timestamp = source["Timestamp"];
	        this.Uri = source["Uri"];
	        this.Msg = source["Msg"];
	    }
	}
	export class Playback {
	    Req: Tmp[];
	    Resp: Tmp[];
	
	    static createFrom(source: any = {}) {
	        return new Playback(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Req = this.convertValues(source["Req"], Tmp);
	        this.Resp = this.convertValues(source["Resp"], Tmp);
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

