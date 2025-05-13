import { Component } from "@angular/core";
import {FormsModule} from "@angular/forms";
     
@Component({
    selector: "file_crypter",
    standalone: true,
    imports: [FormsModule],
    template: `<label>Введите имя:</label>
                 <input [(ngModel)]="name" placeholder="name">
                 <h1>Добро пожаловать {{name}}!</h1>`
})
export class AppComponent { 
    name= "";
}