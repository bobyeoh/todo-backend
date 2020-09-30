docker run -it --name todo-backend \
-p 0.0.0.0:8888:8888 \
-e ENV="dev" \
-e SESSION_EFFECTIVE_DURATION="24h" \
-e COOKIE_KEY="token" \
-e MAIL_SERVER="" \
-e MAIL_PORT="" \
-e MAIL_USERNAME="" \
-e MAIL_PASSWORD="" \
-d todo-backend