import { Component } from "@angular/core";
import {FormsModule} from "@angular/forms";
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatButtonModule} from '@angular/material/button';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
import {HttpClient, HttpClientModule, HttpHandler } from "@angular/common/http";
import { DomSanitizer } from '@angular/platform-browser';
import { buffer } from "rxjs";
     
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
    resultIsDone = false;
    file: any;
    resultFileURL;
    filenameDownload = "";

    constructor(
        private http: HttpClient,
        private sanitizer: DomSanitizer
    ) {}

    uploadFile(event) {
        this.file = event.target.files[0];
        this.filepath = this.file.name;
    }

    reader(file, callback) {
        const fileReader = new FileReader();
        fileReader.onload = () => callback(null, fileReader.result);
        fileReader.onerror = (err) => callback(err);
        fileReader.readAsText(file);
    }

    doAction(): void {
        // this.resultIsDone = !this.resultIsDone
        if (this.action == "encryption") {
            this.reader(this.file, (err, result) => {
                console.log(result);
                this.http.post(`crypter/encrypt?data=${result}`, null).subscribe({next:(data:any) => {
                    console.log(data);
                    const blob = new Blob([data], { type: 'application/octet-stream' });
                    this.resultFileURL = this.sanitizer.bypassSecurityTrustResourceUrl(window.URL.createObjectURL(blob));
                    this.filenameDownload = "encrypted.txt";
                    },
                    error: error => console.log(error)});
            });
        } else {
            this.reader(this.file, (err, result) => {
               const uintArray = new Uint8Array(result);
                this.http.post(`crypter/decrypt?data=${result}`, null).subscribe({next:(data:any) => {
                    console.log(data);
                    const blob = new Blob([data], { type: 'application/octet-stream' });
                    this.resultFileURL = this.sanitizer.bypassSecurityTrustResourceUrl(window.URL.createObjectURL(blob));
                    this.filenameDownload = "decrypted.txt";
                },
                error: error => console.log(error)});
            });
        }
    }
}