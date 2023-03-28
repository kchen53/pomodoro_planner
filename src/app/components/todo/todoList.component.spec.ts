import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TodoListComponent } from './todoList.components';
import { AppModule } from '../../app.module';

describe('TodoListComponent', () => {
  let component: TodoListComponent;
  let fixture: ComponentFixture<TodoListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
      ],
      declarations: [ TodoListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TodoListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('check if addTodo function was called', () => {
    expect(component.addTodo).toHaveBeenCalled;
  });

  it('check if getTodo function was called', () => {
    expect(component.getTodos).toHaveBeenCalled;
  });

  it('check if deleteTodo function was called', () => {
    expect(component.deleteTodo).toHaveBeenCalled;
  });

  // it('check the to-do list html header', () => {
  //   fixture.detectChanges();
  //   const compiled = fixture.nativeElement;
  //   expect(compiled.querySelector('mat-card-title').textContent).toContain('To-Do List');
  // });

  //New Unit Test

});