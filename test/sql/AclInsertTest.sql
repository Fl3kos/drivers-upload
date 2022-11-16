INSERT INTO acl.ecommerce_user4application_role (application_role_code, ecommerce_user_code)
VALUES ((SELECT role_code from acl.application_role where role_name = 'ROLE_APPTMS_DRIVER'),'B0000011'),
((SELECT role_code from acl.application_role where role_name = 'ROLE_APPTMS_DRIVER'),'K0000021');