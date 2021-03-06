---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustom"
description: |-
  Configure custom Internet Services.
---

# fortios_firewall_internetservicecustom
Configure custom Internet Services.

## Argument Reference

The following arguments are supported:

* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service database and custom database.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry.
* `dst` - Destination address or address group name.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block supports:

* `name` - Select the destination address or address group object from available options.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceCustom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicecustom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
