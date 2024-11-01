Using Environment Variables in Nginx Config File

1. Overview
   Sometimes, we need our server to run in different modes and use different resources. For instance, we might need to
   use different resources for testing and running our server in production. Therefore, using environment variables is a
   sound solution to these kinds of problems.

nginx.conf contains config related to our Nginx server. The Nginx config doesn’t support environment variables out of
the box. However, there are tools and workarounds that we can use to carry out this process for us.

In this article, we’ll discuss how we can use environment variables in nginx.conf. Additionally, our approach will apply
to running the server in a container and directly.

2. Using envsubst
   envsubst is a tool that we can use to generate an Nginx config file dynamically before the server starts. It replaces
   variables in a text with the corresponding values of our shell variables.

Moreover, it’s not a standalone utility. It comes with the gettext metapackage. So, by default, it will already be
available on our Linux machine.
https://www.baeldung.com/linux/nginx-config-environment-variables