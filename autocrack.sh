#!/bin/sh

declare -a interfaces
declare -a networks
declare -a clients

interfaces=($(ip link | awk '/state UP/{print $2a;getline}'))
interface
network
client
ap
channel
deauths

printf "\x1b[38;5;4m Looking for up network interfaces...\n"

printf "\x1b[38;5;11m"
for element in "${!interfaces[*]}"; do
	echo "${interfaces[$element]}" | sed 's/://'
done

printf "\x1b[38;5;13m"
printf "choose interface\n"
select word in "${interfaces[*]}"; do
    if [[ -z "$word" ]]; then
        printf '"%s" is not a valid choice\n' "$REPLY" >&2
    else
        user_in="$((REPLY - 1))"
		interface="$word" 
		#DEBUG
        echo "$interface"
        break
    fi
done

interface="${interface}" | awk '{ print substr( $0, 1, length($0)-1 ) }'
#airmon-ng check kill
#airmon-ng start $interface 

interface="${interface}mon"
#DEBUG
echo "$interface"

#timeout 5 airodump-ng $interface
