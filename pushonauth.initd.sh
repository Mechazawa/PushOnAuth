#! /bin/sh
# /etc/init.d/pushonauth


PIDFILE=/var/run/pushonauth.pid

case $1 in
  start)
    echo "Starting pushonauth"
    start-stop-daemon --start --exec /usr/bin/pushonauth -m --pidfile $PIDFILE --background -d /etc/
    ;;
  stop)
    echo "Stopping pushonauth"
    start-stop-daemon --stop --pidfile $PIDFILE
    ;;
  *)
    echo "Usage: /etc/init.d/pushonauth {start|stop}"
    exit 1
    ;;
esac

exit 0
