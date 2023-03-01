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

  ngOnInit() {
    this.getTodos();
  }

  getTodos(): void {
    this.todoService.getTodos()
      .subscribe(todos => this.todos = todos);
  }

  addTodo(task: string): void {
    task = task.trim();
    if (!task) { return; }
    const todo: Todo = { task, complete: false } as Todo;
    this.todoService.addTodo(todo).subscribe(newTodo => {
        this.todos.push(newTodo);
      });
  }

  deleteTodo(todo: Todo): void {
    this.todos = this.todos.filter(t => t !== todo);
    this.todoService.deleteTodo(todo.id).subscribe();
  }

  update(todo: Todo): void {
  }
}
   
