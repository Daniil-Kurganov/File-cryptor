import { Component } from "@angular/core";
import {FormsModule} from "@angular/forms";
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatButtonModule} from '@angular/material/button';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
     
@Component({
    selector: "file_crypter",
    standalone: true,
    imports: [
        FormsModule,
        MatCardModule,
        MatFormFieldModule,
        MatSelectModule,
        MatButtonModule,
        MatButtonToggleModule
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

    uploadFile(event) {
        this.file = event.target.files[0];
        this.filepath = this.file.name;
    }

    doAction(): void {
        this.resultIsDone = !this.resultIsDone
        let fileReader = new FileReader();
        fileReader.onload = (e) => {
            console.log(fileReader.result);
        }
        // fileReader.readAsText(this.file);
        // let s = fileReader.result;
        // console.log(s)

        fileReader.readAsArrayBuffer(this.file);
        let bytes = fileReader.result;
        console.log(bytes);
    }
}