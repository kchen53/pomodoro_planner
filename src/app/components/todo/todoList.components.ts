import { Component, OnInit } from '@angular/core';
import { Todo } from './todo';
import { TodoService } from './todoService';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todoList.component.html',
  styleUrls: ['./todoList.component.css']
})

export class TodoListComponent implements OnInit {

  newTodo: string = '';
  todos: Todo[] = [];

  constructor(private todoService: TodoService) { }

  onEnter() {
    this.newTodo = '';
  }

  //to perform a one-time initialization of tasks
  ngOnInit() {
    this.getTodos();
  }

  getTodos(): void { 
    this.todoService.getTodos() //call to the getTodos() method of the todoService object returns an Observable
      .subscribe(todos => this.todos = todos); //call to the subscribe method of the obervable and takes the array of todos as its argument
  }

  addTodo(name: string): void {
    name = name.trim();
    if (!name) { return; }
    const todo: Todo = { name, complete: false } as Todo;
    this.todoService.addTodo(todo).subscribe(newTodo => {
        this.todos.push(newTodo);
      });
      this.newTodo = '';
  }

  // addTodo(title: string): void {
  //   title = title.trim();
  //   if (!title) { return; }
  //   const todo: Todo = { title, complete: false } as Todo;
  //   this.todoService.addTodo(todo).subscribe(newTodo => {
  //       this.todos.push(newTodo);
  //     });
  //   this.newTodo = '';
  // }

  deleteTodo(todo: Todo): void {
    this.todos = this.todos.filter(t => t !== todo);
    this.todoService.deleteTodo(todo.id).subscribe();
  }

  // deleteTodo(todo: Todo): void {
  //   this.todos = this.todos.filter(t => t !== todo);
  //   this.todoService.deleteTodo(todo.userId).subscribe();
  // }

  update(todo: Todo): void {
  }
}
   
