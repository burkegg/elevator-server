# -*- coding: utf-8 -*-

import unittest

from .. import Vator

class TestVator(unittest.TestCase):

    def test_floor_count(self):
        vator = Vator(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'], 3)
        self.assertEqual(10, vator.floor_count())

    def _get_fid_by_name(self, elevator, floor_name):
        # This is just a helper function to get the floor_id by the floor name
        for fid in elevator.floor_map.keys():
            if elevator.floor_map[fid] == floor_name:
                return fid

    def test_find_closet_car(self):
        # the method Elevator.find_closet_car should take a floor_id (fid) and
        # return the car_id (cid) of the closest car
        vator = Vator(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'], 3)
        # Set up the test so that one car is on the top floor, once car is on
        # floor E and the last car is on the first floor
        top_car, mid_car, bottom_car = list(vator.car_map.keys())
        top_fid = self._get_fid_by_name(vator, 'J')
        mid_fid = self._get_fid_by_name(vator, 'E')
        vator.car_current_floor[top_car] = top_fid
        vator.car_current_floor[mid_car] = mid_fid

        # If a car is on that floor, it should return that car's cid
        self.assertEqual(top_car, vator.find_closest_car(top_fid))
        self.assertEqual(mid_car, vator.find_closest_car(mid_fid))
        self.assertEqual(bottom_car, vator.find_closest_car(vator.first_floor))

        # Unambiguous cases
        fid_h = self._get_fid_by_name(vator, 'H')
        fid_b = self._get_fid_by_name(vator, 'B')
        fid_f = self._get_fid_by_name(vator, 'F')
        self.assertEqual(vator.find_closest_car(fid_h), top_car)
        self.assertEqual(vator.find_closest_car(fid_f), mid_car)
        self.assertEqual(vator.find_closest_car(fid_b), bottom_car)

        # Return one but not both car ids if the cars are equidistant
        fid_c = self._get_fid_by_name(vator, 'C')
        closest_to_c = vator.find_closest_car(fid_c)
        self.assertTrue(closest_to_c == mid_car or closest_to_c == bottom_car)

    def test_call_car(self):
        # the method Elevator.call_Car should take a floor_id (fid) and
        # move the closest car to that floor
        vator = Vator(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'], 2)
        # Both cars are on first floor
        car_one, car_two = list(vator.car_map.keys())
        self.assertEqual(vator.car_current_floor[car_one], vator.first_floor)
        self.assertEqual(vator.car_current_floor[car_two], vator.first_floor)
        # Put car one on top floor
        vator.car_current_floor[car_one] = self._get_fid_by_name(vator, 'J')
        # Call car_one to floor 'H'
        fid_h = self._get_fid_by_name(vator, 'H')
        vator.call_car(fid_h)
        self.assertEqual(vator.car_current_floor[car_one], fid_h)
        # Calling again shouldn't bring the other car to floor F
        vator.call_car(fid_h)
        self.assertEqual(vator.car_current_floor[car_one], fid_h)
        self.assertNotEqual(vator.car_current_floor[car_two], fid_h)
        # Call car_two to floor 'C'
        fid_c = self._get_fid_by_name(vator, 'C')
        vator.call_car(fid_c)
        self.assertEqual(vator.car_current_floor[car_two], fid_c)
        # F is closer to H than to C
        fid_f = self._get_fid_by_name(vator, 'F')
        vator.call_car(fid_f)
        self.assertEqual(vator.car_current_floor[car_one], fid_f)
