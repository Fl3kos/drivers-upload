package sql

import "fmt"

func GenerateAclInsert(allUsers []string, role string) string {
	query := "INSERT INTO acl.ecommerce_user4application_role (application_role_code, ecommerce_user_code)\nVALUES "
	value := "((SELECT role_code from acl.application_role where role_name = '%v'),'%v')"

	for i, user := range allUsers {
		if user != "" {
			valueF := fmt.Sprintf(value, role, user)

			if i != len(allUsers)-1 {
				valueF = valueF + ",\n"
			}

			query = query + valueF
		}

	}

	query = query + ";"

	return query
}

func GenerateAclRoleInsert(usernames []string, shopcode, environoment string) string {
	query := "INSERT INTO `scope-users`.scope_ecommerce_user (ecommerce_user_code, application_code, segmentation) \nVALUES "
	value := "('%v','%v','[{\"type\": \"AND\",\"content\": [{\"values\": [\"%v\"], \"dimension\": \"storeCode\"}]}]')"
	for i, user := range usernames {
		valueF := fmt.Sprintf(value, user, environoment, shopcode)

		if i != len(usernames)-1 {
			valueF = valueF + ",\n"
		}
		query = query + valueF
	}

	query = query + ";"

	return query
}

func GenerateExpeditionLayoutSql(warehouseCode, typeE, locationZone, area, position, templateArea, templatePosition, shippingSecuence, priority, closer_sorter, active, locationTemplate, location string) string {
	query := "INSERT INTO wms.location_expedition ( warehouse_code,type,location_zone,area,position,template_area,template_position,shipping_sequence,priority,closer_sorter,active,location_template,location) VALUES ('%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::integer,'%v'::integer,'%v'::character varying,'%v'::boolean,'%v'::character varying,'%v'::character varying) ON CONFLICT ON CONSTRAINT location_expedition_pkey DO UPDATE SET warehouse_code = EXCLUDED.warehouse_code,type = EXCLUDED.type,location_zone = EXCLUDED.location_zone,area = EXCLUDED.area,position = EXCLUDED.position,template_area = EXCLUDED.template_area,template_position = EXCLUDED.template_position,shipping_sequence = EXCLUDED.shipping_sequence,priority = EXCLUDED.priority,closer_sorter = EXCLUDED.closer_sorter,active = EXCLUDED.active,location_template = EXCLUDED.location_template,location = EXCLUDED.location;"
	fQuery := fmt.Sprintf(query, warehouseCode, typeE, locationZone, area, position, templateArea, templatePosition, shippingSecuence, priority, closer_sorter, active, locationTemplate, location)

	return fQuery
}

func GeneratePickingLayoutSql(warehouse_code, typeP, location_format, template, corridor, module, shelf, gap, location_zone, picking_zone, weight, height, width, length, capacity_fee, restocking_fee, picking_sequence, putaway_sequence, direction, blocked, active, rotation, cycle_count, location_template, location string) string {
	query := "INSERT INTO wms.location_picking ( warehouse_code,type,location_format,template,corridor,module,shelf,gap,location_zone,picking_zone,weight,height,width,length,capacity_fee,restocking_fee,picking_sequence,putaway_sequence,direction,blocked,active,rotation,cycle_count,location_template,location) VALUES ('%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::character varying,'%v'::real,'%v'::real,'%v'::real,'%v'::real,'%v'::real,'%v'::real,'%v'::integer,'%v'::integer,'%v'::character varying,'%v'::boolean,'%v'::boolean,'%v'::character varying,'%v'::boolean,'%v'::character varying,'%v'::character varying) ON CONFLICT ON CONSTRAINT location_picking_pkey DO UPDATE SET warehouse_code = EXCLUDED.warehouse_code,type = EXCLUDED.type,location_format = EXCLUDED.location_format,template = EXCLUDED.template,corridor = EXCLUDED.corridor,module = EXCLUDED.module,shelf = EXCLUDED.shelf,gap = EXCLUDED.gap,location_zone = EXCLUDED.location_zone,picking_zone = EXCLUDED.picking_zone,weight = EXCLUDED.weight,height = EXCLUDED.height,width = EXCLUDED.width,length = EXCLUDED.length,capacity_fee = EXCLUDED.capacity_fee,restocking_fee = EXCLUDED.restocking_fee,picking_sequence = EXCLUDED.picking_sequence,putaway_sequence = EXCLUDED.putaway_sequence,direction = EXCLUDED.direction,blocked = EXCLUDED.blocked,active = EXCLUDED.active,rotation = EXCLUDED.rotation,cycle_count = EXCLUDED.cycle_count,location_template = EXCLUDED.location_template,location = EXCLUDED.location;"
	fQuery := fmt.Sprintf(query, warehouse_code, typeP, location_format, template, corridor, module, shelf, gap, location_zone, picking_zone, weight, height, width, length, capacity_fee, restocking_fee, picking_sequence, putaway_sequence, direction, blocked, active, rotation, cycle_count, location_template, location)

	return fQuery
}
