export namespace model {
	
	export class CreateBlankTemp {
	    type: string;
	    title: string;
	    content: string;
	    tags: string;
	    matters: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateBlankTemp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.tags = source["tags"];
	        this.matters = source["matters"];
	    }
	}

}

