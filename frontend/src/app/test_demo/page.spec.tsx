import { render, screen, fireEvent } from "@testing-library/react";
import TestDemo from "./page";

describe("Counter", () => {
  test("increments count by 1 on button click", async () => {
    render(<TestDemo />);
    expect(screen.queryByText(/Count: 0/)).toBeInTheDocument();
    const button = screen.getByText("+");
    fireEvent.click(button);
    expect(screen.queryByText(/Count: 1/)).toBeInTheDocument();
  });
});
