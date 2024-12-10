export namespace models {
	
	export class PasswordEntry {
	    Id: string;
	    Website: string;
	    Username: string;
	    Timestamp: number;
	    Password: string;
	
	    static createFrom(source: any = {}) {
	        return new PasswordEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Website = source["Website"];
	        this.Username = source["Username"];
	        this.Timestamp = source["Timestamp"];
	        this.Password = source["Password"];
	    }
	}

}

