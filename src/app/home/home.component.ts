import { Component, OnInit, Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{

  public getJsonValue: any;
  public postJsonValue: any;

  constructor(private http: HttpClient){
    
  }
  // posts: any[] = [];
  ngOnInit(): void {
    this.getMethod();
    this.postMethod();
    this.deleteMethod();
  }

  // loadPosts(){
  //   this.http.get('https://jsonplaceholder.typicode.com/posts').subscribe((posts: any[]) => {
  //     this.posts = posts;
  //   });
  // }

  public getMethod(){
    this.http.get('https://jsonplaceholder.typicode.com/todos').subscribe((getJsonValue) => {
      //console.log(data);
      this.getJsonValue = getJsonValue;
    }
    );
  }

  public postMethod(){
    const header = new HttpHeaders({
      contentType: 'application/json'
    })

    let body = {
      id: 'id',
      task: 'task',
      due: 'due',
      complete: 'complete',
    };

    this.http.post('https://jsonplaceholder.typicode.com/todos',body, {headers:header}).subscribe((data) => {
      console.log(data);
      this.postJsonValue = data;
    }
    );
  }

  public deleteMethod(){
    this.http.delete('https://jsonplaceholder.typicode.com/posts/1').subscribe();
  }

}
