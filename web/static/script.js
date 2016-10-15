$( document ).ready(function() {

    // Submits the form when button is clicked
    $('#sbmit').click(submit_form);
    // Submits the form when enter key is pressed
    $('#site').bind("enterKey", submit_form);
    $('#site').keyup(function(e) {
      if(e.keyCode == 13) {
        $(this).trigger("enterKey");
      }
    });

    // Attach the loader to Ajax starts & stops
    var $loader = $('#loader');
    var $info = $('#info');
    var $results = $('#results');
    $(document)
      .ajaxStart(function () {
        $loader.show();
        $results.hide();
        $info.hide();
      })
      .ajaxStop(function () {
        $loader.hide();
        $results.show();
      });

      function submit_form() {
        var api = "/api/check";
        var site_data = $('#site').val();
        var comprehensive_data = false;
        if($('#comprehensive').is(":checked")) {
          comprehensive_data = true;
        }

        $.getJSON( api, {
          domain: site_data,
          comprehensive: comprehensive_data
        }).done(function(data) {
          render_results(data);
        }).fail(function( jqxhr, textStatus, error ) {
          render_error(site_data);
        });
      }

    // Renders results on page according to the JSON data received
    function render_results(jsonData) {
      var data = jsonData;
      var root_domain = data['domain'];
      var valid = data['valid'];
      var weaknesses = data['checks'];

      $('#root-domain').text(root_domain);
      if(!valid) {
        $('#results-invalid').show();
        $('#results-valid').hide();
        return;
      } else {
        $('#results-invalid').hide();
        $('#results-valid').show();
      }

      // clear table
      $('#results-valid').html('');

      $.each(weaknesses, function(index, weakness) {
        var table = '<table><col width="50%"><col width="50%"><tbody>';
        var main_desc = style(weakness['title']);
        var ok_desc = weakness['okDescription'];
        var ok_urls = weakness['okUrls'];
        var not_ok_desc = weakness['notOkDescription'];
        var not_ok_urls = weakness['notOkUrls'];

        table += '<tr><td colspan="2" class="scan-desc">' + main_desc + '</td></tr>';
        table += '<tr><th style="text-align:left">Scanned Domains/URLs:</th><th>Scan Results:</th><tr>';

        table += '<tr>';
        table += '<td class="result"><b>' + not_ok_urls.join('<br>') + '</b></td>';
        table += '<td><div class="oracle reason reason-danger">' + not_ok_desc + '</div></td>'
        table += '</tr>';

        table += '<tr>';
        table += '<td class="result"><b>' + ok_urls.join('<br>') + '</b></td>';
        table += '<td><div class="oracle reason reason-safe">' + ok_desc + '</div></td>'
        table += '</tr>';

        table += '</tbody></table>';
        $('#results-valid').append(table);
      });
    }


    // Displays an error message for invalid inputs
    function render_error(domain) {
      $('#root-domain').text(domain);
      $('#results-invalid').show();
      $('#results-valid').hide();
    }

    function wrap_paragraph(text) {
      return '<p>' + text + '</p>';
    }

    function style(text) {
      var bold_re = /\*([A-z -_\(\)]+)\*/g;
      var bolded_text = text.replace(bold_re, "<b>$1<\/b>");
      var split_bolded_text = bolded_text.split('\n')
      var mapped_bolded_text = split_bolded_text.map(wrap_paragraph);
      return mapped_bolded_text.join('');
    }
});
