#!/bin/bash

# This is for create and setup the Network Namespaces

# 1) Create eath_base and orb_sat Namespaces
ip netns add earth_base
ip netns add orb_sat

# 2) Turn on the loopback interface
ip netns exec earth_base ip link set lo up
ip netns exec orb_sat ip link set lo up

# 3) Create and connect Virtual Ethernet
ip link add veth_earth type veth peer name veth_sat
ip link set veth_earth netns earth_base
ip link set veth_sat netns orb_sat

# 3.1) Turn on
ip netns exec earth_base ip link set veth_earth up
ip netns exec orb_sat ip link set veth_sat up

# 4) Assign IP Addresses
ip netns exec earth_base ip addr add 10.0.1.1/24 dev veth_earth
ip netns exec orb_sat ip addr add 10.0.1.2/24 dev veth_sat

