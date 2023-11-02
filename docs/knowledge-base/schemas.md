# Schemas

#### The page includes all necessary information to better understand the schemas of the zedcloud terraform provider, which can help to prepare a configuraton that could work without any unexpected behavior.

---

### Objects re-creation

Starting from the version 2.2.5-beta the `name` field of the various objects was marked with `ForceNew`. It means that if the name of a object is changed the terraform provider will try to re-create the object. The list of objects that can be re-created are the following:

- application
- application_instance
- datasource
- image
- network
- network_instance
- volume_instance
- node

CATUTION: please, be careful, the action of object re-creatione is a DESTRUCTIVE one. It means that if you decided to change the name of an object and apply the changes, the terraform provider will send a request to delete the object and its relations and after that
a new object with new ralations will be created. `terraform plan` will ONLY show that the object to be recreated will be deleted, but will not show the objects associated with the recreated object as objects to be deleted.

To process an object re-creation, the system will perform the following steps:

- delete the current object\*
- create a new object with the new name

\* It's worth mentioning that if an object couln't be deleted by some reason such as: any relation with other objects exists that prevents the object from deletion, the object couldn't be deleted due to some invariants that cannot be broken, etc, then the object won't be deleted and accordingly a new one won't be created.

The process of object re-creation is little bit tricky if the configuration has got defined both resource of the object and its data. In that case the following steps need to be performed:

- set a new value for the name of the desired object's resource
- `terraform apply` - after executing this command an error may occur saying that the system unable to read the object's data: "The requested resource was not found on this server". It's ok because we haven't set a new value for the object's data in the configuration.
- set a new value for the name of the object's data.
- `terraform apply` - to refresh the state of the object

**Example**:

Step 0: a configuration before the name change

```javascript
resource "zedcloud_application_instance"  "test_tf_provider" {
    name = "old_tf_provider"
    ...
}

data "zedcloud_application_instance"  "test_tf_provider" {
    name = "old_tf_provider"
    ...
}
```

Step 1: change the name field in the object's resource

```javascript
resource "zedcloud_application_instance"  "test_tf_provider" {
    name = "new_tf_provider"
    ...
}

data "zedcloud_application_instance"  "test_tf_provider" {
    name = "old_tf_provider" <-- THIS field stays unchanged!
    ...
}
```

Step 2: apply the changed configuration

```bash
terrraform apply
```

Step 3: change the name field in the object's data

```javascript
resource "zedcloud_application_instance"  "test_tf_provider" {
    name = "new_tf_provider"
    ...
}

data "zedcloud_application_instance"  "test_tf_provider" {
    name = "new_tf_provider"
    ...
}
```

Step 3: apply the object's new state

```bash
terrraform apply
```

---

### Required fields

Some of the fields of different objects were changed from `optional` to `required` or vise versa. The reason of that is to synchronize the terraform provider schemas with the zedcloud enpoint schemas. The following tables shows the changes:

**Application**

| Field               | Previous definition | Current definition |
| ------------------- | ------------------- | ------------------ |
| title               | Optional            | Required           |
| manifest.ac_kind    | Optional            | Required           |
| manifest.ac_version | Optional            | Required           |
| manifest.name       | Optional            | Required           |

**Application instance**

| Field     | Previous definition | Current definition |
| --------- | ------------------- | ------------------ |
| name      | Optional            | Required           |
| title     | Optional            | Required           |
| app_id    | Optional            | Required           |
| device_id | Optional            | Required           |

**Datastore**

| Field   | Previous definition | Current definition |
| ------- | ------------------- | ------------------ |
| title   | Optional            | Required           |
| ds_fqdn | Optional            | Required           |
| ds_type | Optional            | Required           |

**Image**

| Field        | Previous definition | Current definition |
| ------------ | ------------------- | ------------------ |
| datastore_id | Optional            | Required           |
| image_arch   | Optional            | Required           |
| image_format | Optional            | Required           |
| title        | Optional            | Required           |

**Network**

| Field | Previous definition | Current definition |
| ----- | ------------------- | ------------------ |
| ip    | Required            | Optional           |

**Network instance**

| Field     | Previous definition | Current definition |
| --------- | ------------------- | ------------------ |
| device_id | Optional            | Required           |
| kind      | Optional            | Required           |
| title     | Optional            | Required           |

**Node (Edge-Node)**

| Field    | Previous definition | Current definition |
| -------- | ------------------- | ------------------ |
| name     | Optional            | Required           |
| model_id | Optional            | Required           |
| title    | Optional            | Required           |

**Project**

| Field | Previous definition | Current definition |
| ----- | ------------------- | ------------------ |
| title | Optional            | Required           |
| type  | Optional            | Required           |

**Volume instance**

| Field     | Previous definition | Current definition |
| --------- | ------------------- | ------------------ |
| name      | Optional            | Required           |
| title     | Optional            | Required           |
| type      | Optional            | Required           |
| device_id | Optional            | Required           |

Note: To not to duplicate required values twice, in resources and in data configuration, the references can be used, example:

```javascript
resource "zedcloud_application_instance"  "test_tf_provider" {
    model_id = "0f16441a-3342-45e5-8c1c-3714c284de51"
    ...
}

data "zedcloud_application_instance"  "test_tf_provider" {
    model_id = zedcloud_application_instance.test_tf_provider.model_id
    ...
}
```

---

### AppInterface

The Application Instance schema contains the field called `interfaces`, which consists of array of `AppInterface` type. The field is marked as `Optional` but not `Required`. Nevertheless, sometimes when the field isn't defined, an error may occur saying that the number of networks in the Bundle doesn't match the number of interfaces in the application instance: "interface count mismatch with Bundle". This means that at least one network is defined in the Bundle and when creating an application instance, the code checks that the number of interfaces in the application instance is equal to or greater than the number of networks in the application.

Example of working configuration:

```javascript

resource "zedcloud_application" "test_tf_provider" {
    name = "test_tf_provider"
    title = "ubuntu-all-ip"
    ...
    manifest {
        ...
        interfaces {
            name = "indirect"
            directattach = false
            privateip = false
            acls {
                matches {
                    type = "protocol"
                    value = "tcp"
                }
                matches {
                    type = "lport"
                    value = "8022"
                }
                actions {
                    portmap = true
                    portmapto {
                        app_port = 22
                    }
                }
            }
            acls {
                matches {
                    type = "host"
                    value = ""
                }
            }
        }
        ...
    }
}


resource "zedcloud_application_instance"  "test_tf_provider" {
    depends_on = [
        zedcloud_application.test_tf_provider,
        zedcloud_edgenode.test_tf_provider,
        zedcloud_network_instance.test_tf_provider
    ]

    name = "test_tf_provider"
    title = "tf test"
    ...
    interfaces {
        directattach = false
        privateip = false
        netinstname = zedcloud_network_instance.test_tf_provider.name
        intfname = "indirect"
		acls {
            actions {
                drop       = false
                limit      = false
                limitburst = 0
                limitrate  = 0
                portmap    = true

                mapparams {
                    port = 22
                }
            }

            matches {
                type  = "protocol"
                value = "tcp"
            }
            matches {
                type  = "lport"
                value = "8022"
            }
		}
		acls {
            matches {
                type = "host"
            }
		}
    }
    ...
}

```
