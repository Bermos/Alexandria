import { Component, OnInit } from '@angular/core';
import { Module, ModulesService } from "./api_client";
import * as module from "module";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Alexandria';
  modules: { [key: string]: Module; } | undefined;

  constructor(private modulesService: ModulesService) {
  }

  ngOnInit(): void {
    this.modulesService.modulesGet().subscribe(
      value => this.modules = value
    )
  }

}
