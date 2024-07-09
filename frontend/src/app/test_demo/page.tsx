"use client";

import React, { useState } from "react";

export default function TestDemo() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div>Count: {count}</div>
      <button onClick={() => setCount(count + 1)}>+</button>
    </>
  );
}
