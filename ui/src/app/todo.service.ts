import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class TodoService {
  constructor(private httpClient: HttpClient) {}

  getTodoList(userEmail: string) {
    return this.httpClient.get(environment.gateway + '/todo/' + userEmail);
  }

  addTodo(userEmail: string, todo: Todo) {
    return this.httpClient.post(environment.gateway + '/todo/' + userEmail, todo);
  }

  completeTodo(userEmail: string, todo: Todo) {
    return this.httpClient.put(environment.gateway + '/todo/' + userEmail, todo);
  }

  deleteTodo(userEmail: string, todo: Todo) {
    return this.httpClient.delete(environment.gateway + '/todo/'+ userEmail + '/' + todo.id);
  }
}

export class Todo {
  id!: string;
  message!: string;
  complete!: boolean;
}