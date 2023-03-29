import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TodoListComponent } from './todoList.components';
import { AppModule } from '../../app.module';
//import { TodoService } from './todoService';

describe('TodoListComponent', () => {
  let component: TodoListComponent;
  let fixture: ComponentFixture<TodoListComponent>;

  //let myToDoService : TodoService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppModule
        //,HttpClientTestingModule
      ],
      //providers: [TodoService],
      declarations: [ TodoListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TodoListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    //myToDoService = TestBed.inject(myToDoService);
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

  //New Unit Test
  it('check the to-do list html header', () => {
      fixture.detectChanges();
      const compiled = fixture.nativeElement;
      expect(compiled.querySelector('h2').textContent).toContain('To-Do List');
    });
    
});