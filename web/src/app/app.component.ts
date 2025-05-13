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
    protocol = "";
    action = ""
    resultIsDone = true;

    doAction(): void {
        this.resultIsDone = !this.resultIsDone
    }
}