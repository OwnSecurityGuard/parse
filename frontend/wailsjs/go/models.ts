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

