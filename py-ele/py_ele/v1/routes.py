# -*- coding: utf-8 -*-

###
### DO NOT CHANGE THIS FILE
### 
### The code is auto generated, your change will be overwritten by 
### code generating.
###
from __future__ import absolute_import

from .api.welcome import Welcome
from .api.car_car_id_floor_current import CarCarIdFloorCurrent
from .api.floor_count import FloorCount
from .api.admin_inventory_pwd import AdminInventoryPwd


routes = [
    dict(resource=Welcome, urls=['/welcome'], endpoint='welcome'),
    dict(resource=CarCarIdFloorCurrent, urls=['/car/<car_id>/floor/current'], endpoint='car_car_id_floor_current'),
    dict(resource=FloorCount, urls=['/floor_count'], endpoint='floor_count'),
    dict(resource=AdminInventoryPwd, urls=['/admin/inventory/<pwd>'], endpoint='admin_inventory_pwd'),
]