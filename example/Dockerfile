FROM public.ecr.aws/lambda/nodejs:14

# copy the package and package-lock files
# docker can use a cached layer if the files have not chaned
COPY package*.json  ${LAMBDA_TASK_ROOT}/

# Install NPM dependencies for function
# docker can use a cached layer if the previous step used cache
RUN npm install

# copy over all the files
COPY . ${LAMBDA_TASK_ROOT}

# this exposes port 8080 to other containers on the network
EXPOSE 8080

# Set the CMD to your handler
# this expects the a file named `app.js` that exports an async function named `handler`
CMD [ "app.handler" ]
