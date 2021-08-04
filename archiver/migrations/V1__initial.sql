CREATE TABLE archive_requests (
  id serial PRIMARY KEY,
  query text NOT NULL, -- url we send for youtube-dl
  error text -- error message for when it fails
);

-- links users to archive requests
CREATE TABLE user_archive_requests (
  user_id integer NOT NULL,
  archive_id integer NOT NULL REFERENCES archive_requests (id),

  PRIMARY KEY (user_id, archive_id)
);

CREATE TABLE videos (
  id serial PRIMARY KEY,
  ytdl_url text NOT NULL, -- for youtube-dl
  downloaded_video_id integer, -- video id on horahora (indicating that we downloaded the video)
  error text -- error message for when it fails
);

-- links videos to archive requests
CREATE TABLE video_archive_requests (
  video_id integer NOT NULL REFERENCES videos (id),
  archive_id integer NOT NULL REFERENCES archive_requests (id),

  PRIMARY KEY (video_id, archive_id)
);
