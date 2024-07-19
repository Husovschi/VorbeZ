# Use an official nginx image as the base image
FROM nginx:alpine

# Copy the front-end files to the nginx html directory
COPY . /usr/share/nginx/html/

# Start nginx when the container launches
CMD ["nginx", "-g", "daemon off;"]
