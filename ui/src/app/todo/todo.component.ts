import { Component, OnInit } from '@angular/core';
import { TodoService, Todo } from '../todo.service';
import { CognitoService } from '../cognito.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css']
})
export class TodoComponent implements OnInit {

  activeTodos: Todo[] = [];
  completedTodos: Todo[] = [];
  todoMessage!: string;
  userEmail!: string;

  constructor(private todoService: TodoService,
              private cognitoService: CognitoService) { }

  ngOnInit() {
    this.cognitoService.getUserEmail().then((userEmail: string) => {
      this.userEmail = userEmail;
      this.getAll();
    });
  }

  getAll() {
    this.todoService.getTodoList(this.userEmail).subscribe((data) => {
      this.activeTodos = (data as Todo[]).filter((a) => !a.complete);
      this.completedTodos = (data as Todo[]).filter((a) => a.complete);
    });
  }

  addTodo() {
    var newTodo : Todo = {
      message: this.todoMessage,
      id: '',
      complete: false
    };

    this.todoService.addTodo(this.userEmail, newTodo).subscribe(() => {
      this.getAll();
      this.todoMessage = '';
    });
  }

  completeTodo(todo: Todo) {
    this.todoService.completeTodo(this.userEmail, todo).subscribe(() => {
      this.getAll();
    });
  }

  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(this.userEmail, todo).subscribe(() => {
      this.getAll();
    })
  }
}
