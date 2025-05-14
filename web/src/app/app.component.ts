import { Component } from "@angular/core";
import {FormsModule} from "@angular/forms";
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatButtonModule} from '@angular/material/button';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
import {HttpClient, HttpClientModule, HttpHandler } from "@angular/common/http";
     
@Component({
    selector: "file_crypter",
    standalone: true,
    imports: [
        FormsModule,
        MatCardModule,
        MatFormFieldModule,
        MatSelectModule,
        MatButtonModule,
        MatButtonToggleModule,
        HttpClientModule
    ],
    templateUrl: `./app.component.html`,
    styleUrl: `./app.component.css`
})

export class AppComponent { 
    // configuration variables
    filepathPlug = "Here will be filepath of selected file";

    protocol = "";
    action = "";
    filepath = this.filepathPlug;
    resultIsDone = true;
    file: any;

    constructor(private http: HttpClient) {}

    uploadFile(event) {
        this.file = event.target.files[0];
        this.filepath = this.file.name;
    }

    reader(file, callback) {
        const fileReader = new FileReader();
        fileReader.onload = () => callback(null, fileReader.result);
        fileReader.onerror = (err) => callback(err);
        fileReader.readAsArrayBuffer(file);
    }

    doAction(): void {
        this.resultIsDone = !this.resultIsDone
        this.reader(this.file, (err, result) => {
            const uintArray = new Uint8Array(result);
            let body = {data: Array.from(uintArray)}
            let response: number[];
            this.http.post("crypter/start", body).subscribe({next:(data:any) => {
                console.log(data);
            },
            error: error => console.log(error)});
            console.log(response);
        });


        // let fileReader = new FileReader();
        // // fileReader.readAsText(this.file);
        // fileReader.readAsArrayBuffer(this.file);
        // let fileNumbers = fileReader.onload = function(): any {
        //     // console.log(fileReader.result);
        //     let fileData = fileReader.result as ArrayBuffer;
        //     const uintArray = new Uint8Array(fileData);
        //    return Array.from(uintArray)};
        // } as number[];
        // let response: number[];
        // let body = {data: fileNumbers}
        // console.log(body);
        // this.http.post("crypter/start", body).subscribe({next:(data:any) => {
        //     console.log(data);
        // },
        // error: error => console.log(error)});
        // console.log(response);
    }
}