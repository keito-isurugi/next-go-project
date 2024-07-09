"use client";

import { useEffect, useState } from "react";

export default function Todos() {
  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState("");
  const [editingId, setEditingId] = useState(null);
  const [editInput, setEditInput] = useState("");

  useEffect(() => {
    getTodos();
  }, []);

  const getTodos = async () => {
    const res = await fetch("http://localhost:8080/todos");
    const result = await res.json();
    setTodos(result);
  };

  const addTodo = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    if (input) {
      try {
        const res = await fetch("http://localhost:8080/todos", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ title: input }),
        });
        if (res.ok) {
          getTodos(); // 新しいTodoを追加した後、最新のTodoリストを取得
          setInput("");
        } else {
          console.error("Failed to add todo");
        }
      } catch (error) {
        console.error("Error adding todo:", error);
      }
    }
  };

  const deleteTodo = async (id: string) => {
    try {
      const res = await fetch(`http://localhost:8080/todos/${id}`, {
        method: "DELETE",
      });
      if (res.ok) {
        getTodos(); // 削除後、最新のTodoリストを取得
      } else {
        console.error("Failed to delete todo");
      }
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  const startEditing = (todo: any) => {
    setEditingId(todo.id);
    setEditInput(todo.title);
  };

  const cancelEditing = () => {
    setEditingId(null);
    setEditInput("");
  };

  const updateTodo = async (id: string) => {
    try {
      const res = await fetch(`http://localhost:8080/todos/${id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title: editInput, done_flag: false }),
      });
      if (res.ok) {
        getTodos();
        setEditingId(null);
        setEditInput("");
      } else {
        console.error("Failed to update todo");
      }
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  const toggleTodo = async (id: string, title: string) => {
    try {
      const res = await fetch(`http://localhost:8080/todos/${id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title: title, doneFlag: true }),
      });
      if (res.ok) {
        getTodos(); // 完了状態を変更後、最新のTodoリストを取得
      } else {
        console.error("Failed to toggle todo completion");
      }
    } catch (error) {
      console.error("Error toggling todo completion:", error);
    }
  };

  return (
    <>
      <div className="relative py-3 sm:max-w-xl sm:mx-auto">
        <div className="absolute inset-0 bg-gradient-to-r from-cyan-400 to-light-blue-500 shadow-lg transform -skew-y-6 sm:skew-y-0 sm:-rotate-6 sm:rounded-3xl"></div>
        <div className="relative px-4 py-10 bg-white shadow-lg sm:rounded-3xl sm:p-20">
          <div className="max-w-md mx-auto">
            <div className="divide-y divide-gray-200">
              <div className="py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
                <h1 className="text-3xl font-extrabold text-center">
                  Todo App
                </h1>
                <form onSubmit={addTodo} className="mt-8 space-y-6">
                  <input
                    type="text"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    className="w-full px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
                    placeholder="Add new todo..."
                  />
                  <button
                    type="submit"
                    className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    Add Todo
                  </button>
                </form>
                <ul className="mt-8 space-y-4">
                  {todos.map((todo: any) => (
                    <li
                      key={todo.id}
                      className="flex items-center justify-between space-x-3"
                    >
                      {editingId === todo.id ? (
                        <div className="flex-1 flex items-center space-x-2">
                          <input
                            type="text"
                            value={editInput}
                            onChange={(e) => setEditInput(e.target.value)}
                            className="flex-1 px-3 py-2 text-gray-700 border rounded-lg focus:outline-none"
                          />
                          <button
                            onClick={() => updateTodo(todo.id)}
                            className="px-2 py-1 text-xs font-medium text-green-600 bg-green-100 rounded-md hover:bg-green-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                          >
                            Save
                          </button>
                          <button
                            onClick={cancelEditing}
                            className="px-2 py-1 text-xs font-medium text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
                          >
                            Cancel
                          </button>
                        </div>
                      ) : (
                        <>
                          <div
                            onClick={() => toggleTodo(todo.id, todo.title)}
                            className={`flex items-center space-x-3 cursor-pointer ${
                              todo.doneFlag ? "line-through text-gray-400" : ""
                            }`}
                          >
                            <input
                              type="checkbox"
                              checked={todo.doneFlag}
                              readOnly
                              className="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                            />
                            <span>{todo.title}</span>
                          </div>
                          <div className="flex items-center space-x-2">
                            <button
                              onClick={() => startEditing(todo)}
                              className="px-2 py-1 text-xs font-medium text-blue-600 bg-blue-100 rounded-md hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                            >
                              Edit
                            </button>
                            <button
                              onClick={() => deleteTodo(todo.id)}
                              className="px-2 py-1 text-xs font-medium text-red-600 bg-red-100 rounded-md hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                            >
                              Delete
                            </button>
                          </div>
                        </>
                      )}
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
