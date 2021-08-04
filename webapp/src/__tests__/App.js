import { BrowserRouter } from "react-router-dom";
import { render as rtlRender, waitFor } from "@testing-library/react";

import App from "../App";

function render(ui, options) {
  function Wrapper(props) {
    return <BrowserRouter {...props} />;
  }
  return rtlRender(ui, { ...options, wrapper: Wrapper });
}

test("renders front page", async () => {
  const { getByText } = render(<App />);
  const headerText = await waitFor(() => getByText(/horahora/i));
  expect(headerText).toBeInTheDocument();
});
