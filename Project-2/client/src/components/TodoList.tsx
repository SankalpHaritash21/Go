import React, { useState, useEffect } from "react";
import axios from "axios";

interface Todo {
  id: number;
  title: string;
}

const TodoList: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [input, setInput] = useState<string>("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await axios.get<Todo[]>("http://localhost:5000/todos");
      setTodos(response.data);
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  const addTodo = async () => {
    try {
      const response = await axios.post<Todo>("http://localhost:5000/todos", {
        title: input,
      });
      setTodos([...todos, response.data]);
      setInput("");
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  return (
    <div className="max-w-md mx-auto my-8 p-4 border border-gray-300 rounded shadow">
      <h1 className="text-2xl mb-4">TODO List</h1>
      <div className="flex mb-4">
        <input
          type="text"
          className="flex-1 border border-gray-300 rounded px-3 py-2 mr-2"
          placeholder="Enter a new todo..."
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button
          className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
          onClick={addTodo}
        >
          Add
        </button>
      </div>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id} className="border-b border-gray-300 py-2">
            {todo.title}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;
