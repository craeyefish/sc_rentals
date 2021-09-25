import datetime

from django.test import TestCase
from django.utils import timezone

from .models import Item


class ItemModelTests(TestCase):

    def test_created_recently_with_future_question(self):
        """
        created_recently() returns False for questions whose create_date
        is in the future.
        """
        time = timezone.now() + datetime.timedelta(days=30)
        future_item = Item(create_date=time)
        self.assertIs(future_item.created_recently(), False)