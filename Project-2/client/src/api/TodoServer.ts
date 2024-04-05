import { Todo } from "../types/type";

const BASE_URL = "http://localhost:9000";

export const getTodos = async (): Promise<Todo[]> => {
  const response = await fetch(`${BASE_URL}/todos`);
  return response.json();
};

export const addTodo = async (todo: Omit<Todo, "id">): Promise<Todo> => {
  const response = await fetch(`${BASE_URL}/todos/add`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(todo),
  });
  return response.json();
};

export const deleteTodo = async (id: number): Promise<void> => {
  await fetch(`${BASE_URL}/todos/delete`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id }),
  });
};

export const updateTodo = async (id: number, todo: Todo): Promise<Todo> => {
  const response = await fetch(`${BASE_URL}/todos/update`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(todo),
  });
  return response.json();
};
