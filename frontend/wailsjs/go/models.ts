export namespace models {
	
	export class PasswordEntry {
	    Id: string;
	    Website: string;
	    Username: string;
	    Password: string;
	    Timestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new PasswordEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Website = source["Website"];
	        this.Username = source["Username"];
	        this.Password = source["Password"];
	        this.Timestamp = source["Timestamp"];
	    }
	}

}

