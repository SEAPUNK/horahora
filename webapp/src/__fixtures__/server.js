import { rest } from "msw";
import { setupServer } from "msw/node";

const handlers = [
  rest.get("/api/home", (req, res, ctx) => {
    return res(
      ctx.json({
        L: {
          UserID: 1,
          Username: "foo",
          ProfilePictureURL: "",
          Rank: 1,
        },
        PaginationData: {},
        Videos: [],
      })
    );
  }),
];

export const server = setupServer(...handlers);
